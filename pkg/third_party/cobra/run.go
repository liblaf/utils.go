package cobra

import (
	"strings"

	"github.com/dsnet/try"
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func RunCommands(cmd *cobra.Command, args []string) (err error) {
	defer try.Handle(&err)
	root := cmd.Root()
	for _, c := range cmd.Commands() {
		pterm.DefaultSection.Println(c.CommandPath())
		root.SetArgs(append(strings.Fields(c.CommandPath())[1:], args...))
		root.Execute()
	}
	return nil
}
