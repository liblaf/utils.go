package exec

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/dsnet/try"
	er "github.com/liblaf/utils.go/pkg/std/errors"
	"mvdan.cc/sh/v3/syntax"
)

type Cmd struct {
	exec.Cmd
}

func Command(name string, arg ...string) *Cmd {
	r := &Cmd{*exec.Command(name, arg...)}
	return r
}

func (c *Cmd) String() string {
	defer try.F(er.FatalOnError)
	var args []string
	for _, arg := range c.Args {
		arg := try.E1(syntax.Quote(arg, syntax.LangBash))
		args = append(args, arg)
	}
	return strings.Join(args, " ")
}

func (c *Cmd) Bind() *Cmd {
	c.Stdin = os.Stdin
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c
}

func (c *Cmd) EnvAppend(k, v string) *Cmd {
	defer try.F(er.FatalOnError)
	c.Env = append(c.Environ(), fmt.Sprintf("%s=%s", k, try.E1(syntax.Quote(v, syntax.LangBash))))
	return c
}
