package install

import (
	"github.com/liblaf/utils.go/cmd/pkg/install/appimage_launcher"
	"github.com/liblaf/utils.go/cmd/pkg/install/brew"
	"github.com/liblaf/utils.go/cmd/pkg/install/cfw"
	"github.com/liblaf/utils.go/cmd/pkg/install/cloudflare_warp"
	"github.com/liblaf/utils.go/cmd/pkg/install/code"
	"github.com/liblaf/utils.go/cmd/pkg/install/conda"
	"github.com/liblaf/utils.go/cmd/pkg/install/edge"
	"github.com/liblaf/utils.go/cmd/pkg/install/grub_theme"
	"github.com/liblaf/utils.go/cmd/pkg/install/typora"
	"github.com/liblaf/utils.go/cmd/pkg/install/zotero"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "install",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(appimage_launcher.RootCmd)
	RootCmd.AddCommand(brew.RootCmd)
	RootCmd.AddCommand(cfw.RootCmd)
	RootCmd.AddCommand(cloudflare_warp.RootCmd)
	RootCmd.AddCommand(code.RootCmd)
	RootCmd.AddCommand(conda.RootCmd)
	RootCmd.AddCommand(edge.RootCmd)
	RootCmd.AddCommand(grub_theme.RootCmd)
	RootCmd.AddCommand(typora.RootCmd)
	RootCmd.AddCommand(zotero.RootCmd)
}
