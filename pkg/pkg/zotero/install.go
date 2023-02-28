package zotero

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/pkg/interactive/download"
	"github.com/liblaf/utils.go/pkg/std/archive"
	o "github.com/liblaf/utils.go/pkg/std/os"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/liblaf/utils.go/pkg/std/os/ubuntu/desktop"
)

func Install() (err error) {
	defer try.Handle(&err)
	tmpDir := try.E1(os.MkdirTemp("", ""))
	defer func() { try.E(o.Remove(tmpDir)) }()
	tarball := filepath.Join(tmpDir, "zotero.tar.bz2")
	try.E(download.Download(tarballURL(), tarball))
	try.E(archive.Extract(tarball, tmpDir))
	prefix := optPrefix()
	try.E(o.Remove(prefix))
	try.E(o.Move(filepath.Join(tmpDir, "Zotero_linux-x86_64"), prefix))
	try.E(ex.Command(filepath.Join(prefix, "set_launcher_icon")).Bind().Run())
	try.E(o.LinkRelative(filepath.Join(prefix, "zotero.desktop"), filepath.Join(desktop.UserApplications(), "zotero.desktop")))
	return nil
}
