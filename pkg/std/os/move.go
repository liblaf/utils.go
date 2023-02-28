package os

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	fm "github.com/liblaf/utils.go/pkg/std/fmt"
	"github.com/pterm/pterm"
)

func Move(src, dst string) (err error) {
	defer try.Handle(&err)
	try.E(os.MkdirAll(filepath.Dir(dst), DirPerm()))
	try.E(os.Rename(src, dst))
	pterm.Success.Println(fm.To("MOVE", src, dst))
	return nil
}
