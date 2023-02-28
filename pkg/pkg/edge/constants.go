package edge

import (
	"path/filepath"

	"github.com/liblaf/utils.go/pkg/std/os/ubuntu"
)

const sources = `### THIS FILE IS AUTOMATICALLY CONFIGURED ###
# You may comment out this entry, but any other modifications may be lost.
deb [arch=amd64] https://packages.microsoft.com/repos/edge/ stable main
`

func sourcesListPath() string {
	return filepath.Join(ubuntu.AptSourcesListD(), "microsoft-edge.list")
}

const ascURL = "https://packages.microsoft.com/keys/microsoft.asc"

func gpgPath() string {
	return filepath.Join(ubuntu.AptTrustedGpgD(), "microsoft-edge.gpg")
}
