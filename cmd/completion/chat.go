package completion

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
	CCmd.AddCommand(chatCmd)

	chatCmd.Flags().StringVarP(&chatContent, "content", "c", "", "指定要发送的文本咨询内容,最长不超过2048个字节,[必须指定]")
	chatCmd.MarkFlagsRequiredTogether("content")
}

var (
	chatCmd = &cobra.Command{
		Use:     "chat",
		Short:   "send chat content to chatgpt",
		Example: "hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("您搜索的内容:", chatContent)
			if len(chatContent) == 0 {
				fmt.Println("请输出聊天内容")
				return
			}

			var res interface{}
			client := pkg.NewChatClient(context.Background())
			err := client.WithPrompt([]string{chatContent}).Send(&res)
			if err != nil {
				fmt.Println(err)
			}
		},
	}
)
