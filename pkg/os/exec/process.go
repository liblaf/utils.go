package exec

import (
	"fmt"
	"strings"

	"github.com/dsnet/try"
	"github.com/emirpasic/gods/stacks/arraystack"
	er "github.com/liblaf/utils.go/pkg/errors"
	"github.com/pterm/pterm"
)

var stack = arraystack.New()

func Start(a ...any) {
	defer try.F(er.FatalOnError)
	stack.Push(fmt.Sprint(a...))
	pterm.Info.Println(strings.TrimSpace(strings.Repeat("+ ", stack.Size())), er.O1(stack.Peek()))
}

func Complete() {
	defer try.F(er.FatalOnError)
	a := er.O1(stack.Pop())
	pterm.Success.Println(strings.TrimSpace(strings.Repeat("+ ", stack.Size()+1)), a)
}
