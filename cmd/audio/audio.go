package audio

import (
	"github.com/spf13/cobra"
)

const (
	defaultModel = "whisper-1" // 默认模型
)

var (
	file     string // 文件
	model    string // gpt模型
	prompt   string // 提示
	language string // Supplying the input language in ISO-639-1 format will improve accuracy and latency. link-at:https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
)

func init() {
	Cmd.AddCommand(transCmd)
	Cmd.AddCommand(transcCmd)
}

var (
	Cmd = &cobra.Command{
		Use:   "audio",
		Short: "learn how to turn audio into text.",
	}
)
