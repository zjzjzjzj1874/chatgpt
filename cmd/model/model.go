// Package model 查询GPT所有模型
package model

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

var (
	Cmd = &cobra.Command{
		Use:   "model",
		Short: "lists the currently available models,",
		Run: func(cmd *cobra.Command, args []string) {
			var resp pkg.ModelResponse
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodGet), pkg.WithUrl(pkg.MODEL_URL))
			if err != nil {
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("Total Model:%d", len(resp.Data))
			for _, item := range resp.Data {
				color.Cyan("ID:%s", item.ID)
			}
		},
	}
)
