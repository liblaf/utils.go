package avatar

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/errors"
	"github.com/liblaf/utils.go/pkg/interactive/download"
	o "github.com/liblaf/utils.go/pkg/os"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "avatar",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		try.E(os.MkdirAll(filepath.Join(prefix, "jpg"), o.DirPerm()))
		try.E(os.MkdirAll(filepath.Join(prefix, "png"), o.DirPerm()))
		for name, url := range urls {
			jpg := filepath.Join(prefix, "jpg", fmt.Sprintf("%s.jpg", name))
			png := filepath.Join(prefix, "png", fmt.Sprintf("%s.png", name))
			try.E(download.Download(url, jpg))
			try.E(ex.Command("magick", "convert", jpg, png).Bind().Run())
		}
		return nil
	},
}

func init() {
	defer try.F(er.FatalOnError)
	RootCmd.Flags().StringP("prefix", "o", filepath.Join(try.E1(os.Getwd()), "avatar"), "")
	RootCmd.MarkFlagDirname("prefix")
}
