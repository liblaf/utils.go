package edge

import (
	"github.com/liblaf/utils.go/pkg/pkg/edge"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "edge",
	Short: "https://www.microsoft.com/en-us/edge/download",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return edge.Install()
	},
}
