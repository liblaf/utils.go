package os

import (
	"github.com/dsnet/try"
	fm "github.com/liblaf/utils.go/pkg/fmt"
	"github.com/otiai10/copy"
	"github.com/pterm/pterm"
)

func Copy(src, dst string) (err error) {
	try.Handle(&err)
	try.E(copy.Copy(src, dst))
	pterm.Success.Println(fm.To("COPY", src, dst))
	return nil
}
