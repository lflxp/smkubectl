package git

import (
	"fmt"
	"strings"

	"log/slog"

	c "github.com/lflxp/smkubectl/completion"
	"github.com/lflxp/smkubectl/completion/git/tree"
	"github.com/lflxp/smkubectl/utils"
)

// https://refactoringguru.cn/design-patterns/chain-of-responsibility/go/example#example-0
const CC = `$HERE -o=jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'`

type Process struct {
	next c.Cmd
}

// 处理命令类型
func (m *Process) Execute(command *c.Command) {
	var err error
	if command.ProcessDone {
		m.next.Execute(command)
		return
	}

	// 搜索命令
	result := c.Search(&tree.Gits, command.Cmds, command.IsLastWorkSpace)
	if result == nil {
		// 如果长度大于3个则直接执行命令
		// k get po -n kube-system
		// k get po
		slog.Debug("命令不存在,执行原始命令", "Raw", command.Raw)
		// 原始命令解析
		command.Result, err = utils.ExecCmdString(command.Raw)
		command.Result = fmt.Sprintf("命令不存在，执行原生命令 %s\n%s\n", command.Raw, command.Result)
		if err != nil {
			command.Result = fmt.Sprintf("%s\n", err.Error())
		}
	}

	if len(result) > 1 {
		// 只展示推荐命令 不执行
		count := 0
		for k, _ := range result {
			if count == 0 {
				command.Result += "Name\n"
				count++
			}
			// fmt.Printf("%s\n", k)
			// 计算剩余命令
			last := strings.Replace(k, command.Cmds[len(command.Cmds)-1], "", 1)
			command.Result += fmt.Sprintf("%s\n", last)
		}
	} else if len(result) == 1 {
		for _, v := range result {
			if v.IsShell {
				cmd := v.Shell
				if len(command.Cmds) <= 1 {
					cmd = command.Raw
				} else {
					switch command.Cmds[1] {
					case "branch":
						if command.Cmds[len(command.Cmds)-1] == "-d" {
							cmd = "git branch"
						}
					}
				}
				fmt.Println("执行命令", cmd)
				err = utils.ExecCmd(cmd)
				if err != nil {
					command.Result = fmt.Sprintf("Error\n%s\n", err.Error())
					continue
				}
			} else {
				// 处理多命令打印
				count := 0
				// fmt.Println(v)
				for k, _ := range v.Children {
					if count == 0 {
						command.Result += "Name\n"
						count++
					}
					// fmt.Printf("%s\n", k)
					// 计算剩余命令
					last := strings.Replace(k, command.Cmds[len(command.Cmds)-1], "", 1)
					command.Result += fmt.Sprintf("%s\n", last)
				}
			}
		}
	}

	// 解析一级命令
	command.ProcessDone = true
	// command.Result = "Process 命令解析完成"
	m.next.Execute(command)
	slog.Debug("Process 命令解析完成,进入下一步")
}

func (m *Process) SetNext(next c.Cmd) {
	m.next = next
}
