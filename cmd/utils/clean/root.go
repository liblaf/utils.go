package clean

import (
	"github.com/liblaf/utils.go/cmd/utils/clean/apt"
	"github.com/liblaf/utils.go/cmd/utils/clean/brew"
	"github.com/liblaf/utils.go/cmd/utils/clean/cache"
	"github.com/liblaf/utils.go/cmd/utils/clean/npm"
	"github.com/liblaf/utils.go/cmd/utils/clean/pip"
	"github.com/liblaf/utils.go/cmd/utils/clean/tldr"
	"github.com/liblaf/utils.go/cmd/utils/clean/tmp"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "clean",
	Args: cobra.NoArgs,

	RunE: co.RunCommands,
}

func init() {
	RootCmd.AddCommand(apt.RootCmd)
	RootCmd.AddCommand(brew.RootCmd)
	RootCmd.AddCommand(cache.RootCmd)
	RootCmd.AddCommand(npm.RootCmd)
	RootCmd.AddCommand(pip.RootCmd)
	RootCmd.AddCommand(tldr.RootCmd)
	RootCmd.AddCommand(tmp.RootCmd)
}
