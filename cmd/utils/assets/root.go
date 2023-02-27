package assets

import (
	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/cmd/utils/assets/avatar"
	"github.com/liblaf/utils.go/cmd/utils/assets/favicon"
	co "github.com/liblaf/utils.go/pkg/cobra"
	o "github.com/liblaf/utils.go/pkg/os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "assets",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(co.RunCommands(cmd, args))
		try.E(o.Copy("favicon/ico/a.ico", "favicon.ico"))
		return nil
	},
}

func init() {
	RootCmd.AddCommand(avatar.RootCmd)
	RootCmd.AddCommand(favicon.RootCmd)
}
