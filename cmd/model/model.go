// Package model 查询GPT所有模型
package model

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

var (
	Cmd = &cobra.Command{
		Use:     "model",
		Short:   "list all models that gpt support",
		Example: "list model",
		Run: func(cmd *cobra.Command, args []string) {
			var resp pkg.ModelResponse
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodGet), pkg.WithUrl(pkg.MODEL_URL))
			if err != nil {
				fmt.Println("New Client Err:", err)
				return
			}

			err = client.Send(&resp)
			if err != nil {
				fmt.Println("Send Err:", err)
				return
			}

			fmt.Println("Total Model:", len(resp.Data))
			for _, item := range resp.Data {
				fmt.Println("ID:", item.ID)
			}
		},
	}
)
