package pkg

import (
	"context"
	"fmt"
	"testing"
)

func TestChatClient_Send(t *testing.T) {
	t.Run("#Chat", func(t *testing.T) {
		var (
			chatContent = "what is golang?"
		)
		var resp ChatResponse
		client, err := NewChatClient(context.Background())
		if err != nil {
			fmt.Println(err)
			return
		}
		err = client.WithPrompt([]string{chatContent}).Send(&resp)
		if err != nil {
			fmt.Println(err)
		}

		for _, choice := range resp.Choices {
			fmt.Println(choice.Message.Content)
		}
	})
}
