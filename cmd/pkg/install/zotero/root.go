package zotero

import (
	"github.com/liblaf/utils.go/pkg/pkg/zotero"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "zotero",
	Short: "https://www.zotero.org/download/",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) error {
		return zotero.Install()
	},
}
