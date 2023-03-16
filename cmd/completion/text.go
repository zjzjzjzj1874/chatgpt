package completion

import (
	"fmt"
	"github.com/spf13/cobra"
)

var (
	textContent string // 咨询内容
)

func init() {
	CCmd.AddCommand(textCmd)

	textCmd.Flags().StringVarP(&textContent, "content", "c", "", "指定要发送的文本咨询内容,最长不超过2048个字节,[必须指定]")
	textCmd.MarkFlagsRequiredTogether("content")
}

var (
	textCmd = &cobra.Command{
		Use:     "text",
		Short:   "send text to chatgpt",
		Example: "text hello",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("您搜索的内容:", textContent)
		},
	}
)
