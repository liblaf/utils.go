package appimage_launcher

import (
	"github.com/liblaf/utils.go/pkg/pkg/appimage_launcher"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "appimage-launcher",
	Short: "https://github.com/TheAssassin/AppImageLauncher/wiki/Install-on-Ubuntu-or-Debian",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return appimage_launcher.Install()
	},
}
