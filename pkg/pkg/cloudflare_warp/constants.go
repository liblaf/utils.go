package cloudflare_warp

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/liblaf/utils.go/pkg/std/os/ubuntu"
)

func sources() (s string, err error) {
	defer try.Handle(&err)
	codename := strings.TrimSpace((string(try.E1(ex.Command("lsb_release", "-cs").Output()))))
	return fmt.Sprintf("deb [arch=amd64] https://pkg.cloudflareclient.com/ %s main\n", codename), nil
}

func sourcesListPath() string {
	return filepath.Join(ubuntu.AptSourcesListD(), "cloudflare-client.list")
}

const ascURL = "https://pkg.cloudflareclient.com/pubkey.gpg"

func gpgPath() string {
	return filepath.Join(ubuntu.AptTrustedGpgD(), "cloudflare-warp-archive-keyring.gpg")
}
