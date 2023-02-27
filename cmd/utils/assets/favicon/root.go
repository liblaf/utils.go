package favicon

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
	Use:  "favicon",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		color := try.E1(cmd.Flags().GetString("color"))
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		try.E(os.MkdirAll(filepath.Join(prefix, "ico"), o.DirPerm()))
		try.E(os.MkdirAll(filepath.Join(prefix, "png"), o.DirPerm()))
		try.E(os.MkdirAll(filepath.Join(prefix, "svg"), o.DirPerm()))
		for letter := 'a'; letter <= 'z'; letter++ {
			ico := filepath.Join(prefix, "ico", fmt.Sprintf("%c.ico", letter))
			png := filepath.Join(prefix, "png", fmt.Sprintf("%c.png", letter))
			svg := filepath.Join(prefix, "svg", fmt.Sprintf("%c.svg", letter))
			try.E(download.Download(svgURL(letter), svg))
			args := []string{"magick",
				"convert",
				"-background",
				"none",
				svg,
				"-fill",
				color,
				"-colorize",
				"100",
				"-resize",
				"512x512",
				"-gravity",
				"center",
				"-extent",
				"512x512",
			}
			try.E(ex.Command(args[0], append(args[1:], png)...).Bind().Run())
			try.E(ex.Command(args[0], append(args[1:], "-resize", "128x128", ico)...).Bind().Run())
		}
		return nil
	},
}

func init() {
	defer try.F(er.FatalOnError)
	RootCmd.Flags().StringP("color", "c", "#48BEF3", "")
	RootCmd.Flags().StringP("prefix", "o", filepath.Join(try.E1(os.Getwd()), "favicon"), "")
	RootCmd.MarkFlagDirname("prefix")
}
