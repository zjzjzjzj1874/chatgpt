package chat

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

var (
	chatContent string // 咨询内容
)

func init() {
	Cmd.Flags().StringVarP(&chatContent, "content", "c", "", "指定要发送的聊天内容,最长不超过2048个字节")
	_ = Cmd.MarkFlagRequired("content")
}

var (
	Cmd = &cobra.Command{
		Use:     "chat",
		Short:   "creates a completion for the chat message",
		Example: "hello!",
		Run: func(cmd *cobra.Command, args []string) {
			if len(chatContent) == 0 {
				fmt.Println("Please input your chat")
				return
			}

			var resp pkg.ChatResponse
			client, err := pkg.NewChatClient(context.Background())
			if err != nil {
				fmt.Println("New Client Err:", err)
				return
			}

			err = client.WithPrompt([]string{chatContent}).Send(&resp)
			if err != nil {
				fmt.Println("Send Chat Err:", err)
				return
			}

			for _, choice := range resp.Choices {
				fmt.Println(choice.Message.Content)
			}
		},
	}
)
