package zotero

import (
	"path/filepath"

	o "github.com/liblaf/utils.go/pkg/std/os"
)

func tarballURL() string {
	return "https://www.zotero.org/download/client/dl?channel=release&platform=linux-x86_64"
}

func optPrefix() string {
	return filepath.Join(o.UserOpt(), "zotero")
}
