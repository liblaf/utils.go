package archive

import (
	"github.com/dsnet/try"
	"github.com/mholt/archiver/v3"
	"github.com/pterm/pterm"
)

func Extract(source, destination string) (err error) {
	defer try.Handle(&err)
	try.E(archiver.Unarchive(source, destination))
	pterm.Success.Println("EXTRACT", source, "->", destination)
	return nil
}
