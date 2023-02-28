package cfw

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/dsnet/try"
	"github.com/google/go-github/v50/github"
	o "github.com/liblaf/utils.go/pkg/std/os"
)

const (
	user = "Fndroid"
	repo = "clash_for_windows_pkg"
)

const (
	iconURL = "https://github.com/Dreamacro/clash/raw/master/docs/logo.png"
)

func tag() (tag string, err error) {
	defer try.Handle(&err)
	c := github.NewClient(nil)
	r, _ := try.E2(c.Repositories.GetLatestRelease(context.Background(), user, repo))
	return r.GetTagName(), nil
}

func tarballURL(tag string) string {
	return fmt.Sprintf("https://github.com/%s/%s/releases/download/%s/Clash.for.Windows-%s-x64-linux.tar.gz", user, repo, tag, tag)
}

func optPrefix() string {
	return filepath.Join(o.UserOpt(), "cfw")
}
