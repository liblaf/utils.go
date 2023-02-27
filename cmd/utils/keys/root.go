package keys

import (
	"github.com/liblaf/utils.go/cmd/utils/keys/_import"
	"github.com/liblaf/utils.go/cmd/utils/keys/export"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "keys",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.PersistentFlags().StringP("prefix", "p", "keys", "")
	RootCmd.MarkPersistentFlagDirname("prefix")

	RootCmd.AddCommand(export.RootCmd)
	RootCmd.AddCommand(_import.RootCmd)
}
