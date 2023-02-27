package os

import (
	"os"

	"github.com/dsnet/try"
	fm "github.com/liblaf/utils.go/pkg/fmt"
	"github.com/pterm/pterm"
)

func Remove(path string) (err error) {
	defer try.Handle(&err)
	try.E(os.RemoveAll(path))
	pterm.Success.Println(fm.On("REMOVE", path))
	return nil
}
