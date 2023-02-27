package edge

import (
	"github.com/liblaf/utils.go/pkg/pkg/edge"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "edge",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return edge.Install()
	},
}
