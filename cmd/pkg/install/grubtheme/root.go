package grubtheme

import (
	"github.com/liblaf/utils.go/pkg/pkg/grubtheme"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "grub-theme",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return grubtheme.Install()
	},
}
