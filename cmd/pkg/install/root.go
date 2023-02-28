package install

import (
	"github.com/liblaf/utils.go/cmd/pkg/install/brew"
	"github.com/liblaf/utils.go/cmd/pkg/install/cfw"
	"github.com/liblaf/utils.go/cmd/pkg/install/code"
	"github.com/liblaf/utils.go/cmd/pkg/install/conda"
	"github.com/liblaf/utils.go/cmd/pkg/install/edge"
	"github.com/liblaf/utils.go/cmd/pkg/install/grubtheme"
	"github.com/liblaf/utils.go/cmd/pkg/install/typora"
	"github.com/liblaf/utils.go/cmd/pkg/install/zotero"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "install",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(brew.RootCmd)
	RootCmd.AddCommand(cfw.RootCmd)
	RootCmd.AddCommand(code.RootCmd)
	RootCmd.AddCommand(conda.RootCmd)
	RootCmd.AddCommand(edge.RootCmd)
	RootCmd.AddCommand(grubtheme.RootCmd)
	RootCmd.AddCommand(typora.RootCmd)
	RootCmd.AddCommand(zotero.RootCmd)
}
