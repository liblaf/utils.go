package exec

import "github.com/dsnet/try"

func (c *Cmd) Run() (err error) {
	defer try.Handle(&err)
	Start(c)
	defer Complete()
	try.E(c.Cmd.Run())
	return nil
}
