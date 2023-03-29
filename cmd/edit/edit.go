package edit

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

const (
	defaultModel = "text-davinci-edit-001" // 默认模型
)

var (
	input       string // 输入
	model       string // gpt模型
	instruction string // 提示
)

func init() {
	Cmd.Flags().StringVarP(&input, "input", "i", "", "The input text to use as a starting point for the edit.")
	Cmd.Flags().StringVarP(&model, "model", "m", defaultModel, "ID of the model to use. You can use the text-davinci-edit-001 or code-davinci-edit-001 model with this endpoint.")
	Cmd.Flags().StringVarP(&instruction, "instruction", "p", "", "The instruction that tells the model how to edit the prompt.")
	Cmd.MarkFlagsRequiredTogether("instruction", "input")
}

var (
	Cmd = &cobra.Command{
		Use:   "edit",
		Short: "given a prompt and an instruction, the model will return an edited version of the prompt.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(input) == 0 {
				color.Red("Please type input with -i")
				return
			}
			if len(instruction) == 0 {
				color.Red("Please input instruction with -p")
				return
			}
			var (
				req = pkg.EditRequest{
					Input:       input,
					Model:       model,
					Instruction: instruction,
				}
				resp pkg.EditResponse
			)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithUrl(pkg.EDIT_URL), pkg.WithBody(req))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Err:%s", err.Error())
				return
			}

			color.White("您的输入:%s", input)
			for idx, choice := range resp.Choices {
				color.Cyan("输出:%d.%s", idx, choice.Text)
			}
		},
	}
)
