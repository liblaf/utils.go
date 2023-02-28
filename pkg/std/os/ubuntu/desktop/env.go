package desktop

import (
	"path/filepath"

	o "github.com/liblaf/utils.go/pkg/std/os"
)

func UserApplications() string {
	return filepath.Join(o.UserShare(), "applications")
}
