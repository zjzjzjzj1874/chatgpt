package pkg

import (
	"github.com/imroc/req/v3"
	"os"

	"sync"
	"time"
)

type Client struct {
	*req.Client
}

const (
	GPT_KEY = "GPT_KEY"
	GPT_URL = "https://api.openai.com/v1/chat/completions"
)

var (
	GPTClient Client
	once      sync.Once
)

func NewGPTClient() *Client {
	once.Do(func() {
		key := os.Getenv(GPT_KEY)
		if len(key) == 0 {
			panic("empty gpt key")
		}
		client := req.C().
			SetTimeout(30 * time.Second).
			SetCommonBearerAuthToken(key).
			SetCommonContentType("application/json; charset=utf-8")
		GPTClient = Client{Client: client}
	})
	return &GPTClient
}
