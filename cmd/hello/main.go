package main

import (
	"github.com/dsnet/try"
	"github.com/liblaf/utils.go/cmd/hello/hello"
	er "github.com/liblaf/utils.go/pkg/errors"
)

func main() {
	defer try.F(er.FatalOnError)
	try.E(hello.RootCmd.Execute())
}
