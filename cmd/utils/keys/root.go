package keys

import (
	"os"
	"path"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/cmd/utils/keys/export"
	"github.com/liblaf/utils.go/cmd/utils/keys/import_"
	er "github.com/liblaf/utils.go/pkg/std/errors"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "keys",
	Args: cobra.NoArgs,
}

func init() {
	defer try.F(er.FatalOnError)

	RootCmd.PersistentFlags().StringP("prefix", "p", path.Join(try.E1(os.Getwd()), "keys"), "")
	RootCmd.MarkPersistentFlagDirname("prefix")

	RootCmd.AddCommand(export.RootCmd)
	RootCmd.AddCommand(import_.RootCmd)
}
