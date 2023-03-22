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


	response like :
		{
			"choices":[
				{
					"finish_reason":"stop",
					"index":0,
					"message":{
						"content":"\n\n您是一个由计算机程序创造的智能语音助手，使用自然语言处理算法来理解和回复用户的问题和话题。",
						"role":"assistant"
					}
				}
			],
			"created":1679472989,
			"id":"chatcmpl-6wnvxvB7i0301",
			"object":"chat.completion",
			"usage":{
				"completion_tokens":43,
				"prompt_tokens":11,
				"total_tokens":54
			}
		}
*/

var defaultChat = Chat{
	Model: "gpt-3.5-turbo", // TODO 模型以参数传入,不传则默认该模型
}

type ChatClient struct {
	context.Context
	*Client

	Messages []string
}

func NewChatClient(ctx context.Context) (*ChatClient, error) {
	c, err := NewClient()
	if err != nil {
		return nil, err
	}
	return &ChatClient{
		Context: ctx,
		Client:  c,
	}, nil
}

func (c *ChatClient) WithPrompt(messages []string) *ChatClient {
	c.Messages = messages
	return c
}

func (c *ChatClient) Send(src interface{}) (err error) {
	if len(c.Messages) == 0 {
		return ErrWithEmptyContent
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
