package edge

import (
	"io"
	"net/http"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	try.E(installSources())
	try.E(installGPG())
	try.E(ex.Command("sudo", "apt", "update").Bind().Run())
	try.E(ex.Command("sudo", "apt", "install", "microsoft-edge-stable").Bind().Run())
	return nil
}

func installSources() (err error) {
	defer try.Handle(&err)
	c := ex.Command("sudo", "tee", sourcesListPath())
	stdin := try.E1(c.StdinPipe())
	go func() {
		defer stdin.Close()
		stdin.Write([]byte(sources))
	}()
	try.E(c.Run())
	return nil
}

func installGPG() (err error) {
	defer try.Handle(&err)
	resp := try.E1(http.Get(ascURL))
	defer resp.Body.Close()
	c := ex.Command("sudo", "gpg", "--dearmor", "--output", trustedGpgPath())
	stdin := try.E1(c.StdinPipe())
	go func() {
		defer stdin.Close()
		io.Copy(stdin, resp.Body)
	}()
	try.E(c.Run())
	return nil
}
