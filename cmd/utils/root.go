package main

import (
	"github.com/liblaf/utils.go/cmd/hello/docs"
	"github.com/liblaf/utils.go/cmd/utils/assets"
	"github.com/liblaf/utils.go/cmd/utils/check_unlock_media"
	"github.com/liblaf/utils.go/cmd/utils/clean"
	"github.com/liblaf/utils.go/cmd/utils/keys"
	"github.com/liblaf/utils.go/cmd/utils/pretty"
	"github.com/liblaf/utils.go/cmd/utils/update"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "utils",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(assets.RootCmd)
	RootCmd.AddCommand(check_unlock_media.RootCmd)
	RootCmd.AddCommand(clean.RootCmd)
	RootCmd.AddCommand(docs.RootCmd)
	RootCmd.AddCommand(keys.RootCmd)
	RootCmd.AddCommand(pretty.RootCmd)
	RootCmd.AddCommand(update.RootCmd)
}
