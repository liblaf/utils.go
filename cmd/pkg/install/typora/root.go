package typora

import (
	"github.com/liblaf/utils.go/pkg/pkg/typora"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "typora",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return typora.Install()
	},
}
