package pretty

import (
	"github.com/liblaf/utils.go/cmd/utils/pretty/_go"
	"github.com/liblaf/utils.go/cmd/utils/pretty/python"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "pretty",
	Args: cobra.NoArgs,

	RunE: co.RunCommands,
}

func init() {
	RootCmd.AddCommand(_go.RootCmd)
	RootCmd.AddCommand(python.RootCmd)
}
