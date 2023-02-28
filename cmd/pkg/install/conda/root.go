package conda

import (
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/pkg/pkg/conda"
	o "github.com/liblaf/utils.go/pkg/std/os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "conda",
	Short: "https://conda.io/projects/conda/en/latest/user-guide/install/macos.html",
	Args:  cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		try.E(conda.Install(conda.Args{
			Batch:  try.E1(cmd.Flags().GetBool("batch")),
			Force:  try.E1(cmd.Flags().GetBool("force")),
			Prefix: try.E1(cmd.Flags().GetString("prefix")),
			Skip:   try.E1(cmd.Flags().GetBool("skip")),
			Update: try.E1(cmd.Flags().GetBool("update")),
			Tests:  try.E1(cmd.Flags().GetBool("tests")),
		}))
		return nil
	},
}

func init() {
	RootCmd.Flags().BoolP("batch", "b", false, "run install in batch mode (without manual intervention), it is expected the license terms (if any) are agreed upon")
	RootCmd.Flags().BoolP("force", "f", false, "no error if install prefix already exists")
	RootCmd.Flags().StringP("prefix", "p", filepath.Join(o.UserOpt(), "conda"), "install prefix, must not contain spaces.")
	RootCmd.Flags().BoolP("skip", "s", false, "skip running pre/post-link/install scripts")
	RootCmd.Flags().BoolP("update", "u", false, "update an existing installation")
	RootCmd.Flags().BoolP("tests", "t", false, "run package tests after installation (may install conda-build)")
}
