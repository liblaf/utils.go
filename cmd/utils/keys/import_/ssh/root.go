package ssh

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	ex "github.com/liblaf/utils.go/pkg/std/os/exec"
	pa "github.com/liblaf/utils.go/pkg/std/path"
	"github.com/spf13/cobra"
)

var keyTypes = []string{"dsa", "ecdsa", "ecdsa_sk", "ed25519", "ed25519_sk", "rsa"}

var RootCmd = &cobra.Command{
	Use:  "ssh",
	Args: cobra.NoArgs,

	RunE: func(cmd *cobra.Command, args []string) (err error) {
		defer try.Handle(&err)
		prefix := try.E1(cmd.Flags().GetString("prefix"))
		home := try.E1(os.UserHomeDir())
		targetDirectory := filepath.Join(home, ".ssh")

		configSrc := filepath.Join(prefix, "ssh", "config")
		if try.E1(pa.Exists(configSrc)) {
			try.E(ex.Command("install", "-D", "--target-directory", targetDirectory, configSrc).Bind().Run())
		}
		for _, keyType := range keyTypes {
			keySrc := filepath.Join(prefix, "ssh", "id_"+keyType)
			pubKeySrc := keySrc + ".pub"
			if try.E1(pa.Exists(keySrc)) {
				try.E(ex.Command("install", "-D", "--target-directory", targetDirectory, keySrc).Bind().Run())
				try.E(ex.Command("install", "-D", "--target-directory", targetDirectory, pubKeySrc).Bind().Run())
			}
		}

		return nil
	},
}
