package image

import (
	"github.com/spf13/cobra"
)

const (
	defaultNum        = 1    // 默认返回的数量
	minNum            = 1    // 最小数量
	maxNum            = 10   // 最大数量
	imageMaxPromptLen = 1000 // prompt长度限制
)

var (
	image       string // 图片路径
	mask        string // 图片路径
	imagePrompt string // 图片描述内容
	num         int    // 生成图片数量
	size        string // 生成图片大小
)

func init() {
	Cmd.AddCommand(cCmd)
	Cmd.AddCommand(eCmd)
	Cmd.AddCommand(vCmd)
}

var (
	Cmd = &cobra.Command{
		Use:   "img",
		Short: "given a prompt and/or an input image, the model will generate a new image.",
	}
)

// 图片尺寸大小
var imgGenSizeMap = map[string]struct{}{
	"1024x1024": {},
	"512x512":   {},
	"256x256":   {},
}
