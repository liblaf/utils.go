package grub_theme

import (
	"github.com/liblaf/utils.go/pkg/pkg/grub_theme"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "grub-theme",
	Short: "https://github.com/vinceliuice/grub2-themes",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return grub_theme.Install()
	},
}
