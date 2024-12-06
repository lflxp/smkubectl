package git

import (
	c "github.com/lflxp/smkubectl/completion"
)

func Start(command *c.Command) {
	raw := c.Raw{}
	process := Process{}
	out := c.Out{}

	raw.SetNext(&process)
	process.SetNext(&out)

	raw.Execute(command)
}
