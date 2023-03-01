package _go

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "go",

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		if len(args) == 0 {
			args = append(args, "./...")
		}
		try.E(gofmt(args...))
		return nil
	},
}

func gofmt(files ...string) (err error) {
	defer try.Handle(&err)
	args := []string{"fmt"}
	args = append(args, files...)
	try.E(ex.Command("go", args...).Bind().Run())
	return nil
}
