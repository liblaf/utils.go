package import_

import (
	"github.com/liblaf/utils.go/cmd/utils/keys/import_/gpg"
	"github.com/liblaf/utils.go/cmd/utils/keys/import_/ssh"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "import",
	Args: cobra.NoArgs,

	RunE: co.RunCommands,
}

func init() {
	RootCmd.AddCommand(gpg.RootCmd)
	RootCmd.AddCommand(ssh.RootCmd)
}
