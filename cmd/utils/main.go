package main

import (
	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/std/errors"
)

func main() {
	defer try.F(er.FatalOnError)
	try.E(RootCmd.Execute())
}
