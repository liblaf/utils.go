package os

import (
	"os"
	"path/filepath"

	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/std/errors"
)

func UserLocal() string {
	defer try.F(er.FatalOnError)
	home := try.E1(os.UserHomeDir())
	return filepath.Join(home, ".local")
}

func UserBin() string {
	return filepath.Join(UserLocal(), "bin")
}

func UserOpt() string {
	return filepath.Join(UserLocal(), "opt")
}

func UserShare() string {
	return filepath.Join(UserLocal(), "share")
}
