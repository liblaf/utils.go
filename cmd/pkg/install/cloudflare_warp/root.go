package cloudflare_warp

import (
	"github.com/liblaf/utils.go/pkg/pkg/cloudflare_warp"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "cloudflare-warp",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return cloudflare_warp.Install()
	},
}
