package file

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

func init() {
	cCmd.Flags().StringVarP(&id, "id", "i", "", "The ID of the file to use for this request.")
	cCmd.MarkFlagsRequiredTogether("id")
}

var (
	cCmd = &cobra.Command{
		Use:   "detail",
		Short: "Returns the contents of the specified file",
		Run: func(cmd *cobra.Command, args []string) {
			var resp pkg.FileContentResponse
			path := fmt.Sprintf("%s/%s/content", pkg.FILE_URL, id)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodGet), pkg.WithUrl(path))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}
			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Err:%s", err.Error())
				return
			}

			//color.Cyan("Total List:%v", len(resp.Data))

			res, _ := json.Marshal(resp)
			color.Cyan(string(res))
		},
	}
)
