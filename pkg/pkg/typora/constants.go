package typora

import (
	"path/filepath"

	"github.com/liblaf/utils.go/pkg/std/os/ubuntu"
)

const ascURL = "https://typora.io/linux/public-key.asc"

func ascPath() string {
	return filepath.Join(ubuntu.AptTrustedGpgD(), "typora.asc")
}
