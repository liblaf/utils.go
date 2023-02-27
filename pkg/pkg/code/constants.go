package code

import (
	"path/filepath"

	"github.com/liblaf/utils.go/pkg/os/ubuntu"
)

const sources = `### THIS FILE IS AUTOMATICALLY CONFIGURED ###
# You may comment out this entry, but any other modifications may be lost.
deb [arch=amd64,arm64,armhf] http://packages.microsoft.com/repos/code stable main
`

func sourcesListPath() string {
	return filepath.Join(ubuntu.AptSourcesListD(), "vscode.list")
}

const ascURL = "https://packages.microsoft.com/keys/microsoft.asc"

func trustedGpgPath() string {
	return filepath.Join(ubuntu.AptTrustedGpgD(), "microsoft.gpg")
}

var extensions = []string{
	"aaron-bond.better-comments",
	"chouzz.vscode-better-align",
	"eamodio.gitlens",
	"esbenp.prettier-vscode",
	"foxundermoon.shell-format",
	"golang.go",
	"James-Yu.latex-workshop",
	"ms-python.isort",
	"ms-python.python",
	"redhat.vscode-yaml",
	"tamasfe.even-better-toml",
	"WakaTime.vscode-wakatime",
}
