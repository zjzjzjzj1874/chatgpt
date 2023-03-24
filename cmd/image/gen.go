package image

import (
	"fmt"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

var (
	imagePrompt string // 图片描述内容
	num         int    // 生成图片数量
	size        string // 生成图片大小
)

func init() {
	Cmd.Flags().StringVarP(&imagePrompt, "prompt", "p", "", "A text description of the desired image(s). The maximum length is 1000 characters")
	Cmd.Flags().IntVarP(&num, "num", "n", 1, "The number of images to generate. Must be between 1 and 10")
	Cmd.Flags().StringVarP(&size, "size", "s", "1024x1024", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
	_ = Cmd.MarkFlagRequired("prompt")
}

var (
	Cmd = &cobra.Command{
		Use:     "img",
		Short:   "creates an image given a prompt.",
		Example: "A cute baby cat",
		Run: func(cmd *cobra.Command, args []string) {
			if len(imagePrompt) == 0 {
				fmt.Println("Please input your prompt")
				return
			}
			if len(imagePrompt) > 1000 {
				fmt.Println("The maximum length is 1000 characters")
				return
			}
			if num < 1 || num > 10 {
				fmt.Println("The number of images to generate. Must be between 1 and 10")
				return
			}
			if _, ok := imgGenSizeMap[size]; !ok {
				fmt.Println("The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
				return
			}

			var (
				req = pkg.ImageCreateRequest{
					Prompt: imagePrompt,
					Num:    num,
					Size:   size,
				}
				resp pkg.ImageCreateResponse
			)
			client, err := pkg.NewClient(pkg.WithMethod(http.MethodPost), pkg.WithUrl(pkg.IMG_CREATE_URL), pkg.WithBody(req))
			if err != nil {
				fmt.Println("New Client Err:", err)
				return
			}

			err = client.Send(&resp)
			if err != nil {
				fmt.Println("Send Err:", err)
				return
			}

			fmt.Println("Total Image:", len(resp.Data))
			for _, item := range resp.Data {
				fmt.Println("Url:", item.URL)
			}
		},
	}
)

// 图片尺寸大小
var imgGenSizeMap = map[string]struct{}{
	"1024x1024": {},
	"512x512":   {},
	"256x256":   {},
}
