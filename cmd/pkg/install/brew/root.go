package brew

import (
	"github.com/liblaf/utils.go/pkg/pkg/brew"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "brew",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		return brew.Install()
	},
}
