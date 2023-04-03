package file

import (
	"bytes"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zjzjzjzj1874/chatgpt/pkg"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

func init() {
	uCmd.Flags().StringVarP(&file, "file", "f", "", "A text description of the desired image(s). The maximum length is 1000 characters")
	uCmd.Flags().StringVarP(&purpose, "purpose", "p", "", "The intended purpose of the uploaded documents.Use \"fine-tune\" for Fine-tuning. This allows us to validate the format of the uploaded file.")
	uCmd.MarkFlagsRequiredTogether("file", "purpose")
}

var (
	uCmd = &cobra.Command{
		Use:   "upload",
		Short: "Upload a file that contains document(s) to be used across various endpoints/features. Currently, the size of all the files uploaded by one organization can be up to 1 GB. Please contact us if you need to increase the storage limit.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(purpose) == 0 {
				color.Red("%s", "Please input your purpose with -p")
				return
			}
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
			_ = writer.WriteField("purpose", purpose)
			_ = writer.Close()

			var (
				resp pkg.FileUploadResponse
			)

			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithContentType(writer.FormDataContentType()), pkg.WithUrl(pkg.FILE_URL), pkg.WithBody(body))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("返回ID:%s", resp.ID)
		},
	}
)
