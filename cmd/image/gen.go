package image

import (
	"net/http"

	"github.com/fatih/color"
	"github.com/spf13/cobra"

	"github.com/zjzjzjzj1874/chatgpt/pkg"
)

const (
	defaultNum        = 1    // 默认返回的数量
	minNum            = 1    // 最小数量
	maxNum            = 10   // 最大数量
	imageMaxPromptLen = 1000 // prompt长度限制
)

var (
	imagePrompt string // 图片描述内容
	num         int    // 生成图片数量
	size        string // 生成图片大小
)

func init() {
	Cmd.Flags().StringVarP(&imagePrompt, "prompt", "p", "", "A text description of the desired image(s). The maximum length is 1000 characters")
	Cmd.Flags().IntVarP(&num, "num", "n", defaultNum, "The number of images to generate. Must be between 1 and 10")
	Cmd.Flags().StringVarP(&size, "size", "s", "256x256", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
	_ = Cmd.MarkFlagRequired("prompt")
}

var (
	Cmd = &cobra.Command{
		Use:     "img",
		Short:   "creates an image given a prompt.",
		Example: "A cute baby cat",
		Run: func(cmd *cobra.Command, args []string) {
			if len(imagePrompt) == 0 {
				color.Red("%s", "Please input your prompt")
				return
			}
			if len(imagePrompt) > imageMaxPromptLen {
				color.Red("The maximum length is %s characters", imageMaxPromptLen)
				return
			}
			if num < minNum || num > maxNum {
				color.Red("The number of images to generate. Must be between %s and %s", minNum, maxNum)
				return
			}
			if _, ok := imgGenSizeMap[size]; !ok {
				color.Red("%s", "The size of the generated images. Must be one of 256x256, 512x512, or 1024x1024")
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
				color.Red("New Client Err:%s", err.Error())
				return
			}

			err = client.Send(&resp)
			if err != nil {
				color.Red("Send Chat Err:%s", err.Error())
				return
			}

			color.Cyan("Total Image:%v", len(resp.Data))
			for _, item := range resp.Data {
				color.Cyan("Url:%s", item.URL)
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
