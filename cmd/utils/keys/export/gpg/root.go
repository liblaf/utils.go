package gpg

import (
	"os"
	"path"

	"github.com/dsnet/try"
	o "github.com/liblaf/utils.go/pkg/std/os"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "gpg",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		try.E(os.MkdirAll(path.Join(prefix, "gpg"), o.DirPerm()))
		try.E(ex.Command("gpg", "--export-secret-keys", "--armor", "--output", path.Join(prefix, "gpg", "secret.asc")).Bind().Run())
		return nil
	},
}
