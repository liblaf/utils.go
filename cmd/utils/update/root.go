package update

import (
	"github.com/liblaf/utils.go/cmd/utils/update/apt"
	"github.com/liblaf/utils.go/cmd/utils/update/brew"
	"github.com/liblaf/utils.go/cmd/utils/update/npm"
	"github.com/liblaf/utils.go/cmd/utils/update/snap"
	"github.com/liblaf/utils.go/cmd/utils/update/tldr"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "update",

	RunE: co.RunCommands,
}

func init() {
	RootCmd.AddCommand(apt.RootCmd)
	RootCmd.AddCommand(brew.RootCmd)
	RootCmd.AddCommand(npm.RootCmd)
	RootCmd.AddCommand(snap.RootCmd)
	RootCmd.AddCommand(tldr.RootCmd)
}
