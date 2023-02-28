package download

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/dsnet/try"
	fm "github.com/liblaf/utils.go/pkg/std/fmt"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/pterm/pterm"
)

type DownloadFunc func(url string, output string) error

func Download(url string, output string) (err error) {
	defer try.Handle(&err)
	try.E(os.MkdirAll(filepath.Dir(output), os.ModePerm))
	var f DownloadFunc
	if _, err := exec.LookPath("https"); err == nil {
		f = httpie
	}
	if f == nil {
		return errors.New("download tool not found")
	}
	try.E(f(url, output))
	pterm.Success.Println(fm.To("DOWNLOAD", url, output))
	return nil
}

func httpie(url string, output string) error {
	var c *ex.Cmd
	if len(output) > 0 {
		c = ex.Command("https", "--body", "--output", output, "--download", url)
	} else {
		c = ex.Command("https", "--body", "--download", url)
	}
	return c.Bind().Run()
}
