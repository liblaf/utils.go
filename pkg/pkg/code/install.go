package code

import (
	"io"
	"net/http"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	try.E(installGPG())
	try.E(addSources())
	try.E(ex.Command("sudo", "apt", "update").Bind().Run())
	try.E(ex.Command("sudo", "apt", "install", "code").Bind().Run())
	try.E(installExtensions())
	return nil
}

func installGPG() (err error) {
	defer try.Handle(&err)
	resp := try.E1(http.Get(ascURL))
	defer resp.Body.Close()
	c := ex.Command("sudo", "gpg", "--dearmor", "--output", gpgPath())
	stdin := try.E1(c.StdinPipe())
	go func() {
		defer stdin.Close()
		io.Copy(stdin, resp.Body)
	}()
	try.E(c.Run())
	return nil
}

func addSources() (err error) {
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

func installExtensions() (err error) {
	defer try.Handle(&err)
	for _, ext := range extensions {
		try.E(ex.Command("code", "--install-extension", ext, "--force").Bind().Run())
	}
	return nil
}
