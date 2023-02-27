package tmp

import (
	"errors"
	"io/fs"
	"path/filepath"

	"github.com/dsnet/try"
	o "github.com/liblaf/utils.go/pkg/os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "tmp",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		for _, p := range try.E1(filepath.Glob("/tmp/*")) {
			if err := o.Remove(p); err != nil {
				switch {
				case errors.Is(err, fs.ErrPermission):
					continue
				default:
					try.E(err)
				}
			}
		}
		return nil
	},
}
