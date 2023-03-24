package chat

import (
	"context"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

var (
	chatContent string // 咨询内容
)

func init() {
	// TODO add role for chat link at :https://platform.openai.com/docs/guides/chat/chat-vs-completions
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
				color.Red("%s", "Please input your chat")
				return
			}

			var resp pkg.ChatResponse
			client, err := pkg.NewChatClient(context.Background())
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.WithPrompt([]string{chatContent}).Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			for _, choice := range resp.Choices {
				color.Cyan(choice.Message.Content)
			}
		},
	}
)
