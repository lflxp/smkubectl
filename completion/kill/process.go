package kill

import (
	"fmt"
	"strings"

	c "github.com/lflxp/smkubectl/completion"
	"github.com/lflxp/smkubectl/completion/kill/tree"
	"github.com/lflxp/smkubectl/utils"
)

type Process struct {
	next c.Cmd
}

func (m *Process) Execute(command *c.Command) {
	var err error
	if command.ProcessDone {
		m.next.Execute(command)
		return
	}

	result := c.Search(&tree.Kill, command.Cmds, command.IsLastWorkSpace)
	if result == nil {
		command.Result, err = utils.ExecCmdString(command.Raw)
		if err != nil {
			command.Result = fmt.Sprintf("Error: %s\n", err.Error())
		}
		m.next.Execute(command)
		return
	}

	if len(result) > 1 {
		count := 0
		for k := range result {
			if count == 0 {
				command.Result += "Name\n"
				count++
			}
			last := strings.Replace(k, command.Cmds[len(command.Cmds)-1], "", 1)
			command.Result += fmt.Sprintf("%s\n", last)
		}
	} else if len(result) == 1 {
		for _, v := range result {
			if v.IsShell && v.Shell != "" {
				command.Result, err = utils.ExecCmdString(v.Shell)
				if err != nil {
					command.Result = fmt.Sprintf("Error: %s\n", err.Error())
				}
			} else if v.Children != nil {
				count := 0
				for k := range v.Children {
					if count == 0 {
						command.Result += "Name\n"
						count++
					}
					last := strings.Replace(k, command.Cmds[len(command.Cmds)-1], "", 1)
					command.Result += fmt.Sprintf("%s\n", last)
				}
			}
		}
	}

	command.ProcessDone = true
	m.next.Execute(command)
}

func (m *Process) SetNext(next c.Cmd) {
	m.next = next
}