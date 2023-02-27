package conda

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/pkg/interactive/download"
	o "github.com/liblaf/utils.go/pkg/os"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
)

type Args struct {
	Batch  bool
	Force  bool
	Prefix string
	Skip   bool
	Update bool
	Tests  bool
}

func (args *Args) AsList() (res []string) {
	if args.Batch {
		res = append(res, "-b")
	}
	if args.Force {
		res = append(res, "-f")
	}
	if len(args.Prefix) > 0 {
		res = append(res, "-p", args.Prefix)
	}
	if args.Skip {
		res = append(res, "-s")
	}
	if args.Update {
		res = append(res, "-u")
	}
	if args.Tests {
		res = append(res, "-t")
	}
	return res
}

func Install(args Args) (err error) {
	defer try.Handle(&err)
	tmpDir := try.E1(os.MkdirTemp("", ""))
	defer o.Remove(tmpDir)
	installerPath := filepath.Join(tmpDir, "conda.sh")
	try.E(download.Download(installerURL(), installerPath))
	c := ex.Command("bash", installerPath)
	c.Args = append(c.Args, args.AsList()...)
	try.E(c.Bind().Run())
	try.E(postIntall(args.Prefix))
	return nil
}

func postIntall(prefix string) (err error) {
	defer try.Handle(&err)
	condaExePath := filepath.Join(prefix, "bin", "conda")
	try.E(ex.Command(condaExePath, "init", "zsh").Bind().Run())
	return nil
}
