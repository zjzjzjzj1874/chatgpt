package utils

import (
	"testing"

	"github.com/fatih/color"
)

func TestChatClient_Send(t *testing.T) {
	t.Run("#Color-Print", func(t *testing.T) {
		color.Cyan("Prints text in cyan.")

		color.Blue("Prints %s in blue.", "text")
		color.Red("Prints %s in red.", "text")

		color.Magenta("And many others ..")

	})
}
