package brew

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	o "github.com/liblaf/utils.go/pkg/os"
	ex "github.com/liblaf/utils.go/pkg/os/exec"
)

func Install() (err error) {
	defer try.Handle(&err)
	tmpDir := try.E1(os.MkdirTemp("", ""))
	defer func() { try.E(o.Remove(tmpDir)) }()
	try.E(ex.Command("git", "clone", "--depth", "1", brewInstallGitRemote, tmpDir).Bind().Run())
	try.E(ex.Command("bash", filepath.Join(tmpDir, "install.sh")).
		EnvAppend("HOMEBREW_BREW_GIT_REMOTE", HOMEBREW_BREW_GIT_REMOTE).
		EnvAppend("HOMEBREW_CORE_GIT_REMOTE", HOMEBREW_CORE_GIT_REMOTE).
		EnvAppend("HOMEBREW_BOTTLE_DOMAIN", HOMEBREW_BOTTLE_DOMAIN).
		Bind().Run())
	return nil
}
