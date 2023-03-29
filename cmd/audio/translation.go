package audio

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

func init() {
	transCmd.Flags().StringVarP(&file, "file", "f", "", "The audio file to translate, in one of these formats: mp3, mp4, mpeg, mpga, m4a, wav, or webm.")
	transCmd.Flags().StringVarP(&model, "model", "m", defaultModel, "ID of the model to use. Only whisper-1 is currently available.")
	transCmd.Flags().StringVarP(&prompt, "prompt", "p", "", "An optional text to guide the model's style or continue a previous audio segment. The prompt should be in English.")
	transCmd.Flags().StringVarP(&language, "language", "l", "", "The language of the input audio. Supplying the input language in ISO-639-1 format will improve accuracy and latency.")
	transCmd.MarkFlagsRequiredTogether("file")
}

var (
	transCmd = &cobra.Command{
		Use:   "trans",
		Short: "Translates audio into into text.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(file) == 0 {
				color.Red("%s", "Please input your file with -f")
				return
			}
			fi, err := os.Open(file)
			if err != nil {
				color.Red("Open file(%s) failure:%s", file, err.Error())
				return
			}

			body := new(bytes.Buffer)
			writer := multipart.NewWriter(body)
			part, err := writer.CreateFormFile("file", file)
			if err != nil {
				color.Red("CreateFormFile file(%s) failure:%s", file, err.Error())
				return
			}
			_, err = io.Copy(part, fi)
			if err != nil {
				color.Red("Copy file(%s) failure:%s", file, err.Error())
				return
			}
			if len(model) != 0 {
				_ = writer.WriteField("model", model)
			}
			if len(prompt) != 0 {
				_ = writer.WriteField("prompt", prompt)
			}
			if len(language) != 0 {
				_ = writer.WriteField("language", language)
			}
			_ = writer.Close()
			var (
				resp pkg.AudioTranslationResponse
			)

			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithContentType(writer.FormDataContentType()), pkg.WithUrl(pkg.AUDIO_TRANSLATION_URL), pkg.WithBody(body))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("翻译结果:%s", resp.Text)
		},
	}
)
