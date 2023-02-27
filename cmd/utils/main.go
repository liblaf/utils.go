package main

import (
	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/errors"
)

func main() {
	defer try.F(er.FatalOnError)
	try.E(RootCmd.Execute())
}
