package python

import (
	"os"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "python",

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		if len(args) == 0 {
			args = append(args, try.E1(os.Getwd()))
		}
		try.E(isort(args...))
		try.E(black(args...))
		return nil
	},
}

func isort(files ...string) (err error) {
	defer try.Handle(&err)
	args := []string{"--profile", "black"}
	args = append(args, files...)
	try.E(ex.Command("isort", args...).Bind().Run())
	return nil
}

func black(files ...string) (err error) {
	defer try.Handle(&err)
	try.E(ex.Command("black", files...).Bind().Run())
	return nil
}
