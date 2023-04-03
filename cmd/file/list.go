package file

import (
	"github.com/fatih/color"
	"github.com/spf13/cobra"
	"github.com/zjzjzjzj1874/chatgpt/pkg"
	"net/http"
)

var (
	id      string // 文件id
	file    string // 文件
	purpose string // 目的
)

var (
	lCmd = &cobra.Command{
		Use:   "list",
		Short: "Returns a list of files that belong to the user's organization.",
		Run: func(cmd *cobra.Command, args []string) {
			var resp pkg.FileListResponse
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodGet), pkg.WithUrl(pkg.FILE_URL))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("Total List:%v", len(resp.Data))
			for _, meta := range resp.Data {
				color.Cyan("File ID:%v", meta.ID)

			}
		},
	}
)
