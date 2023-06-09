package main

import (
	"context"
	"os"

	"github.com/fatih/color"
	"github.com/zjzjzjzj1874/chatgpt/cmd/audio"
	"github.com/zjzjzjzj1874/chatgpt/cmd/chat"
	"github.com/zjzjzjzj1874/chatgpt/cmd/edit"
	"github.com/zjzjzjzj1874/chatgpt/cmd/embed"
	"github.com/zjzjzjzj1874/chatgpt/cmd/file"
	"github.com/zjzjzjzj1874/chatgpt/cmd/image"
	"github.com/zjzjzjzj1874/chatgpt/cmd/model"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:     "gptx",
		Short:   "gptx is a command-line tool for call openai api",
		Version: version,
	}
)

func init() {
	rootCmd.AddCommand(model.Cmd)
	rootCmd.AddCommand(chat.Cmd)
	rootCmd.AddCommand(image.Cmd)
	rootCmd.AddCommand(audio.Cmd)
	rootCmd.AddCommand(edit.Cmd)
	rootCmd.AddCommand(embed.Cmd)
	rootCmd.AddCommand(file.Cmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		color.Red("RootCmd exec failure:[err:%s]", err)
		os.Exit(1)
	}
}
