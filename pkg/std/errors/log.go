package errors

import (
	"github.com/pterm/pterm"
)

func FatalOnError(a ...any) {
	pterm.Fatal.PrintOnError(a...)
}
