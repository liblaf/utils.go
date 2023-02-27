package npm

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "npm",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(ex.Command("pnpm", "update", "--global").Bind().Run())
		return nil
	},
}
