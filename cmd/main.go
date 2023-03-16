package main

import (
	"context"
	"fmt"
	"github.com/zjzjzjzj1874/chatgpt/cmd/completion"
	"os"

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
	rootCmd.AddCommand(completion.CCmd)
	rootCmd.CompletionOptions.DisableDefaultCmd = true
}

func main() {
	ctx := context.Background()

	if err := rootCmd.ExecuteContext(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
