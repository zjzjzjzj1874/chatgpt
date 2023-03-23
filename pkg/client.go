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
	clientTimeoutSec int64  // http请求超时时间(单位:s)
	prompt           string // 请求体
	method           string // 请求方法
	url              string // 请求url
}

type Option func(client *Client)

func WithPrompt(prompt string) Option {
	return func(c *Client) {
		c.prompt = prompt
	}
}

func WithMethod(method string) Option {
	return func(c *Client) {
		c.method = method
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
}

func NewClient(opts ...Option) (client *Client, err error) {
	key := os.Getenv(GPT_KEY)
	if len(key) == 0 {
		return nil, errors.New("please set your gpt key")
	}

	for _, opt := range opts {
		opt(client)
	}

	client.PreNewClient()
	client = &Client{
		Client: req.C().
			SetTimeout(time.Duration(client.clientTimeoutSec) * time.Second).
			SetCommonBearerAuthToken(key).
			SetCommonContentType("application/json; charset=utf-8")}

	return
}

func (c *Client) Send(src interface{}) (err error) {
	_, err = c.R().SetSuccessResult(src).Send(c.method, c.url)
	return
}
