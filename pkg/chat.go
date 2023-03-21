package pkg

import (
	"context"
)

/*
curl call like:

	curl -X POST https://api.openai.com/v1/chat/completions \
	-H "Content-Type: application/json" \
	-H "Authorization: Bearer $OPENAI_API_KEY" \
	-d '{
	"model": "gpt-3.5-turbo",
	"messages": [{"role": "user", "content": "Hello!"}]
	}'
*/

var defaultChat = Chat{
	Model: "gpt-3.5-turbo",
}

type ChatClient struct {
	context.Context
	*Client

	Messages []string
}

func NewChatClient(ctx context.Context) *ChatClient {
	return &ChatClient{
		Context: ctx,
		Client:  NewClient(),
	}
}

func (c *ChatClient) WithPrompt(messages []string) *ChatClient {
	c.Messages = messages
	return c
}

func (c *ChatClient) Send(src interface{}) (err error) {
	if len(c.Messages) == 0 {
		return nil
	}

	msgs := make([]Message, 0, len(c.Messages))
	for _, message := range c.Messages {
		msgs = append(msgs, Message{
			Role:    "user",
			Content: message,
		})
	}
	defaultChat.Messages = msgs

	return c.R().SetBody(defaultChat).SetSuccessResult(src).MustPost(GPT_URL).Err
}
