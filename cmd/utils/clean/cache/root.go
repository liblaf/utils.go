package cache

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	o "github.com/liblaf/utils.go/pkg/os"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "cache",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		for _, p := range try.E1(pathes()) {
			try.E(o.Remove(p))
		}
		return nil
	},
}

func pathes() (rtn []string, err error) {
	defer try.Handle(&err)
	for _, p := range try.E1(patterns()) {
		newPathes := try.E1(filepath.Glob(p))
		rtn = append(rtn, newPathes...)
	}
	return
}

func patterns() (rtn []string, err error) {
	defer try.Handle(&err)
	cache := try.E1(os.UserCacheDir())
	home := try.E1(os.UserHomeDir())
	rtn = []string{
		cache,
		filepath.Join(home, ".profile"),
		filepath.Join(home, "*bash*"),
		filepath.Join(home, "*zcompdump*"),
	}
	return rtn, nil
}
