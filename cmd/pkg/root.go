package main

import (
	"github.com/liblaf/utils.go/cmd/hello/docs"
	"github.com/liblaf/utils.go/cmd/hello/hello"
	"github.com/liblaf/utils.go/cmd/pkg/install"
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:  "pkg",
	Args: cobra.NoArgs,
}

func init() {
	RootCmd.AddCommand(docs.RootCmd)
	RootCmd.AddCommand(hello.RootCmd)
	RootCmd.AddCommand(install.RootCmd)
}
