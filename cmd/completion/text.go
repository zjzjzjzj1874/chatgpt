package completion

import (
	"github.com/spf13/cobra"
)

func init() {
	CCmd.AddCommand(textCmd)

	textCmd.Flags().StringVarP(&TextContent, "content", "c", "", "指定要发送的文本咨询内容,最长不超过2048个字节,[必须指定]")
	textCmd.MarkFlagsRequiredTogether("content")
}

var (
	TextContent string // 咨询内容
	textCmd     = &cobra.Command{
		Use:                        "text",
		Aliases:                    nil,
		SuggestFor:                 nil,
		Short:                      "send text to chatgpt",
		GroupID:                    "",
		Long:                       "",
		Example:                    "",
		ValidArgs:                  nil,
		ValidArgsFunction:          nil,
		Args:                       nil,
		ArgAliases:                 nil,
		BashCompletionFunction:     "",
		Deprecated:                 "",
		Annotations:                nil,
		Version:                    "",
		PersistentPreRun:           nil,
		PersistentPreRunE:          nil,
		PreRun:                     nil,
		PreRunE:                    nil,
		Run:                        nil,
		RunE:                       nil,
		PostRun:                    nil,
		PostRunE:                   nil,
		PersistentPostRun:          nil,
		PersistentPostRunE:         nil,
		FParseErrWhitelist:         cobra.FParseErrWhitelist{},
		CompletionOptions:          cobra.CompletionOptions{},
		TraverseChildren:           false,
		Hidden:                     false,
		SilenceErrors:              false,
		SilenceUsage:               false,
		DisableFlagParsing:         false,
		DisableAutoGenTag:          false,
		DisableFlagsInUseLine:      false,
		DisableSuggestions:         false,
		SuggestionsMinimumDistance: 0,
	}
)
