package export

import (
	"github.com/liblaf/utils.go/cmd/utils/keys/export/gpg"
	"github.com/liblaf/utils.go/cmd/utils/keys/export/ssh"
	co "github.com/liblaf/utils.go/pkg/third_party/cobra"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "export",
	Args: cobra.NoArgs,

	RunE: co.RunCommands,
}

func init() {
	RootCmd.AddCommand(gpg.RootCmd)
	RootCmd.AddCommand(ssh.RootCmd)
}
