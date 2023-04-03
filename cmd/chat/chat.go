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
	Cmd.Flags().StringVarP(&chatContent, "content", "c", "", "Creates a completion for the chat message")
	_ = Cmd.MarkFlagRequired("content")
}

var (
	Cmd = &cobra.Command{
		Use:   "chat",
		Short: "given a chat conversation, the model will return a chat completion response.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(chatContent) == 0 {
				color.Red("%s", "Please input your chat with -c")
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
			//
			//res, _ := json.Marshal(resp)
			//fmt.Println(string(res))

			for _, choice := range resp.Choices {
				color.Cyan(choice.Message.Content)
			}
		},
	}
)
