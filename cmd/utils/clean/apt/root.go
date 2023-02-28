package apt

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "apt",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(ex.Command("sudo", "apt", "autoremove").Bind().Run())
		return nil
	},
}
