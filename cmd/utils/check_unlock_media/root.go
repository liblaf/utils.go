package check_unlock_media

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/pkg/interactive/download"
	o "github.com/liblaf/utils.go/pkg/std/os"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use: "check-unlock-media",

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		tmpDir := try.E1(os.MkdirTemp("", ""))
		defer o.Remove(tmpDir)
		scriptPath := filepath.Join(tmpDir, "check.sh")
		try.E(download.Download(scriptURL, scriptPath))
		try.E(ex.Command("bash", scriptPath).Bind().Run())
		try.E(o.Remove("bahamut_cookie.txt"))
		return nil
	},
}
