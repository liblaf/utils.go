package pip

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "pip",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(ex.Command("conda", "clean", "--all").Bind().Run())
		try.E(ex.Command("pip", "cache", "purge").Bind().Run())
		return nil
	},
}
