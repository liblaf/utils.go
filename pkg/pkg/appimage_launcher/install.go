package appimage_launcher

import (
	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	try.E(addAptRepository())
	try.E(ex.Command("sudo", "apt", "update").Bind().Run())
	try.E(ex.Command("sudo", "apt", "install", "appimagelauncher").Bind().Run())
	return nil
}

func addAptRepository() (err error) {
	defer try.Handle(&err)
	try.E(ex.Command("sudo", "add-apt-repository", "ppa:appimagelauncher-team/stable").Bind().Run())
	return nil
}
