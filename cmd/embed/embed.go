package embed

import (
	"encoding/json"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

const (
	defaultModel = "text-embedding-ada-002" // 默认模型
)

var (
	input string // 输入
	model string // gpt模型
)

func init() {
	Cmd.Flags().StringVarP(&input, "input", "i", "", "Input text to get embeddings for, encoded as a string or array of tokens. To get embeddings for multiple inputs in a single request, pass an array of strings or array of token arrays. Each input must not exceed 8192 tokens in length.")
	Cmd.Flags().StringVarP(&model, "model", "m", defaultModel, "ID of the model to use. You can use the List models API to see all of your available models, or see our Model overview for descriptions of them.")
	Cmd.MarkFlagsRequiredTogether("input")
}

var (
	Cmd = &cobra.Command{
		Use:   "embed",
		Short: "Creates an embedding vector representing the input text.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(input) == 0 {
				color.Red("Please type input with -i")
				return
			}
			var (
				req = pkg.EmbedRequest{
					Input: input,
					Model: model,
				}
				resp pkg.EmbedResponse
			)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithUrl(pkg.EMBED_URL), pkg.WithBody(req))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Err:%s", err.Error())
				return
			}

			color.White("您的输入:%v", input)
			//color.White("输出:%v", resp)
			res, _ := json.Marshal(resp)
			color.White("输出:%v", string(res))
		},
	}
)
