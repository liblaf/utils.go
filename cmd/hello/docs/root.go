package docs

import (
	"os"

	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/std/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var RootCmd = &cobra.Command{
	Use:  "docs",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		doc.GenMarkdownTree(cmd.Root(), prefix)
		return nil
	},
}

func init() {
	defer try.F(er.FatalOnError)
	RootCmd.Flags().StringP("prefix", "p", try.E1(os.Getwd()), "")
	RootCmd.MarkFlagDirname("prefix")
}
