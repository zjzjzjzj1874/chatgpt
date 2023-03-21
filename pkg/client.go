package pkg

import (
	"github.com/imroc/req/v3"
	"os"

	"time"
)

type GptClient interface {
	Send(interface{}) error
}

type Client struct {
	*req.Client
	prompt string // 请求体
	url    string // 请求url
}

type Option func(client *Client)

func WithPrompt(prompt string) Option {
	return func(c *Client) {
		c.prompt = prompt
	}
}

func WithUrl(url string) Option {
	return func(c *Client) {
		c.url = url
	}
}

const (
	GPT_KEY = "GPT_KEY" // 以环境变量来存放gpt的key
	GPT_URL = "https://api.openai.com/v1/chat/completions"
)

func NewClient(opts ...Option) (client *Client) {
	key := os.Getenv(GPT_KEY)
	if len(key) == 0 {
		panic("empty gpt key")
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
	return c.R().SetSuccessResult(src).MustPost(c.url).Err
}
