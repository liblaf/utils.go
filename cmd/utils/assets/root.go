package assets

import (
	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/cmd/utils/assets/avatar"
	"github.com/liblaf/utils.go/cmd/utils/assets/favicon"
	o "github.com/liblaf/utils.go/pkg/std/os"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
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
