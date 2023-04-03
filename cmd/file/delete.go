package file

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zjzjzjzj1874/chatgpt/pkg"
	"net/http"
)

func init() {
	dCmd.Flags().StringVarP(&id, "id", "i", "", "The ID of the file to use for this request.")
	dCmd.MarkFlagsRequiredTogether("id")
}

var (
	dCmd = &cobra.Command{
		Use:   "del",
		Short: "Delete a file.",
		Run: func(cmd *cobra.Command, args []string) {
			if len(id) == 0 {
				color.Red("%s", "Please input your id with -i")
				return
			}
			var (
				resp pkg.FileDeleteResponse
			)

			path := fmt.Sprintf("%s/%s", pkg.FILE_URL, id)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodDelete), pkg.WithUrl(path))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Err:%s", err.Error())
				return
			}

			color.Cyan("%+v", resp)
		},
	}
)
