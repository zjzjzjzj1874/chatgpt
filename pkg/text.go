package pkg

import "context"

var defaultTextClient = Text2Cmd{
	Model:            "text-davinci-003",
	Temperature:      0,
	MaxTokens:        0,
	TopP:             1.0,
	FrequencyPenalty: 0.2,
	PresencePenalty:  0,
	Stop:             []string{"\n"},
}

type TextClient struct {
	context.Context
	*Client
}

func NewTextClient(ctx context.Context) *TextClient {
	return &TextClient{
		Context: ctx,
		Client:  NewClient(),
	}
}

func (c *TextClient) WithPrompt(prompt string) *TextClient {
	c.prompt = prompt
	return c
}

func (c *TextClient) Send(src interface{}) (err error) {
	defaultTextClient.Prompt = c.prompt
	return c.R().SetBody(defaultTextClient).SetSuccessResult(src).MustPost(GPT_URL).Err
}
