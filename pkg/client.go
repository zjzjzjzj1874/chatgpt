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
	prompt string // 请求体
	method string // 请求方法
	url    string // 请求url
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

func WithUrl(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

const (
	GPT_KEY = "GPT_KEY" // 以环境变量来存放gpt的key

	GPT_URL   = "https://api.openai.com/v1/chat/completions" // POST&GET:和gpt进行聊天
	MODEL_URL = "https://api.openai.com/v1/models"           // GET:请求模型列表
)

func NewClient(opts ...Option) (client *Client, err error) {
	key := os.Getenv(GPT_KEY)
	if len(key) == 0 {
		return nil, errors.New("please set your gpt key")
	}

	client = &Client{
		Client: req.C().
			SetTimeout(30 * time.Second).
			SetCommonBearerAuthToken(key).
			SetCommonContentType("application/json; charset=utf-8")}

	for _, opt := range opts {
		opt(client)
	}

	return
}

func (c *Client) Send(src interface{}) (err error) {
	_, err = c.R().SetSuccessResult(src).Send(c.method, c.url)
	return
}
