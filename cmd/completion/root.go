package completion

import (
	"github.com/spf13/cobra"
)

var CCmd = &cobra.Command{
	Use:   "api",
	Short: "send api to gpt",
}
