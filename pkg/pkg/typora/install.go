package typora

import (
	"io"
	"net/http"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	try.E(installGPG())
	try.E(addAptRepository())
	try.E(ex.Command("sudo", "apt", "update").Bind().Run())
	try.E(ex.Command("sudo", "apt", "install", "typora").Bind().Run())
	return nil
}

func installGPG() (err error) {
	defer try.Handle(&err)
	resp := try.E1(http.Get(ascURL))
	defer resp.Body.Close()
	c := ex.Command("sudo", "tee", ascPath())
	stdin := try.E1(c.StdinPipe())
	go func() {
		defer stdin.Close()
		io.Copy(stdin, resp.Body)
	}()
	try.E(c.Run())
	return nil
}

func addAptRepository() (err error) {
	defer try.Handle(&err)
	try.E(ex.Command("sudo", "add-apt-repository", "deb https://typora.io/linux ./").Bind().Run())
	return nil
}
