package cfw

import (
	"github.com/liblaf/utils.go/pkg/pkg/cfw"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "cfw",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return cfw.Install()
	},
}
