package os

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	fm "github.com/liblaf/utils.go/pkg/std/fmt"
	"github.com/pterm/pterm"
)

func Link(oldname, newname string) (err error) {
	defer try.Handle(&err)
	try.E(os.MkdirAll(filepath.Dir(newname), DirPerm()))
	try.E(os.Remove(newname))
	try.E(os.Symlink(oldname, newname))
	pterm.Success.Println(fm.To("LINK", newname, oldname))
	return nil
}

func LinkRelative(oldname, newname string) (err error) {
	defer try.Handle(&err)
	try.E(Link(try.E1(filepath.Rel(filepath.Dir(newname), oldname)), newname))
	return nil
}
