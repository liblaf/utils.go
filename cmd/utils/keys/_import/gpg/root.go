package gpg

import (
	"path"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "gpg",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		try.E(ex.Command("gpg", "--import", path.Join(prefix, "gpg", "secret.asc")).Bind().Run())
		return nil
	},
}
