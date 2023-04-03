package file

import "github.com/spf13/cobra"

func init() {
	Cmd.AddCommand(lCmd)
	Cmd.AddCommand(uCmd)
	Cmd.AddCommand(dCmd)
	Cmd.AddCommand(cCmd)
}

var (
	Cmd = &cobra.Command{
		Use:   "file",
		Short: "Files are used to upload documents that can be used with features like Fine-tuning.",
	}
)
