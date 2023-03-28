package pkg

import (
	"errors"
	"os"
	"time"

	"github.com/imroc/req/v3"
)

type GptClient interface {
	Send(interface{}) error
}

type Client struct {
	*req.Client
	clientTimeoutSec int64       // http请求超时时间(单位:s)
	prompt           string      // 请求体
	method           string      // 请求方法
	url              string      // 请求url
	body             interface{} // 请求body
	contentType      string      // 类型
}

type Option func(client *Client)

func WithPrompt(prompt string) Option {
	return func(c *Client) {
		c.prompt = prompt
	}
}

func WithContentType(contentType string) Option {
	return func(c *Client) {
		c.contentType = contentType
	}
}

func WithMethod(method string) Option {
	return func(c *Client) {
		c.method = method
	}
}

func WithBody(body interface{}) Option {
	return func(c *Client) {
		c.body = body
	}
}

func WithClientTimeoutSec(timeoutSec int64) Option {
	return func(c *Client) {
		c.clientTimeoutSec = timeoutSec
	}
}

func WithUrl(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

// PreNewClient 预检参数
func (c *Client) PreNewClient() {
	if c.clientTimeoutSec <= 0 {
		c.clientTimeoutSec = default_timeout
	}
	if c.contentType == "" {
	}
}

// PostClient 后置处理参数
func (c *Client) PostClient() {
	if c.contentType != "" {
		c.Client = c.Client.SetCommonContentType(c.contentType)
	}
}

func NewClient(opts ...Option) (client *Client, err error) {
	key := os.Getenv(GPT_KEY)
	if len(key) == 0 {
		return nil, errors.New("please set your gpt key")
	}

	client = &Client{}
	for _, opt := range opts {
		opt(client)
	}

	client.PreNewClient()
	client.Client = req.C().
		SetTimeout(time.Duration(client.clientTimeoutSec) * time.Second).
		SetCommonBearerAuthToken(key).
		SetCommonContentType("application/json; charset=utf-8")
	client.PostClient()
	return
}

func (c *Client) Send(src interface{}) (err error) {
	request := c.R()

	if c.body != nil {
		request = request.SetBody(c.body)
	}
	respErr := ResponseErr{}
	resp, err := request.SetSuccessResult(src).SetErrorResult(&respErr).Send(c.method, c.url)
	if resp.IsErrorState() {
		return respErr.Error
	}

	// TODO add a debug var to print blow info
	//color.Cyan("Resp:%v", src)
	//res, _ := json.Marshal(src)
	//color.Cyan("Total Res:%v", string(res))
	return
}
