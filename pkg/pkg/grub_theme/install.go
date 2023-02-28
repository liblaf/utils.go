package grub_theme

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	o "github.com/liblaf/utils.go/pkg/std/os"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	tmpDir := try.E1(os.MkdirTemp("", ""))
	defer func() { try.E(o.Remove(tmpDir)) }()
	try.E(ex.Command("git", "clone", "--depth", "1", gitRemote, tmpDir).Bind().Run())
	installer := filepath.Join(tmpDir, "install.sh")
	try.E(ex.Command("sed", "--in-place", "s/clear/:/g", installer).Bind().Run())
	try.E(ex.Command("sudo", "bash", installer, "--screen", "4k").Bind().Run())
	return nil
}
