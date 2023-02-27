package cfw

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/pkg/archive"
	"github.com/liblaf/utils.go/pkg/interactive/download"
	o "github.com/liblaf/utils.go/pkg/os"
	"github.com/liblaf/utils.go/pkg/os/ubuntu/desktop"
)

func Install() (err error) {
	defer try.Handle(&err)
	tag := try.E1(tag())
	url := tarballURL(tag)
	tmpDir := try.E1(os.MkdirTemp("", ""))
	defer func() { try.E(o.Remove(tmpDir)) }()
	output := filepath.Join(tmpDir, "cfw.tar.gz")
	try.E(download.Download(url, output))
	try.E(archive.Extract(output, tmpDir))
	prefix := optPrefix()
	try.E(o.Remove(prefix))
	try.E(o.Move(filepath.Join(tmpDir, fmt.Sprintf("Clash for Windows-%s-x64-linux", tag)), prefix))
	try.E(o.LinkRelative(filepath.Join(prefix, "cfw"), filepath.Join(o.UserBin(), "cfw")))
	try.E(installDesktop(prefix))
	return nil
}

func installDesktop(prefix string) (err error) {
	defer try.Handle(&err)
	iconPath := filepath.Join(prefix, "logo.png")
	try.E(download.Download(iconURL, iconPath))
	e := desktop.New("Application", "Clash for Windows")
	e.Icon = iconPath
	e.Exec = filepath.Join(prefix, "cfw")
	try.E(e.Install())
	return nil
}
