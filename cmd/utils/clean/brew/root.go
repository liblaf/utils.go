package brew

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "brew",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(ex.Command("brew", "autoremove").Bind().Run())
		try.E(ex.Command("brew", "cleanup").Bind().Run())
		return nil
	},
}
