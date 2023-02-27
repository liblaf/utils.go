package main

import (
	"github.com/liblaf/utils.go/cmd/utils/assets"
	"github.com/liblaf/utils.go/cmd/utils/clean"
	"github.com/liblaf/utils.go/cmd/utils/keys"
	"github.com/liblaf/utils.go/cmd/utils/update"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "utils",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(assets.RootCmd)
	RootCmd.AddCommand(clean.RootCmd)
	RootCmd.AddCommand(keys.RootCmd)
	RootCmd.AddCommand(update.RootCmd)
}
