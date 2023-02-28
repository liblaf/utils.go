package code

import (
	"github.com/liblaf/utils.go/pkg/pkg/code"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "code",
	Short: "https://code.visualstudio.com",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return code.Install()
	},
}
