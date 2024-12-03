package kubectl

import (
	"fmt"

	"log/slog"

	"github.com/lflxp/smkubectl/completion/kubectl/tree"
)

// https://refactoringguru.cn/design-patterns/chain-of-responsibility/go/example#example-0

type Process struct {
	next Cmd
}

// 处理命令类型
func (m *Process) execute(command *Command) {
	var err error
	if command.ProcessDone {
		m.next.execute(command)
		return
	}

	// 搜索命令
	result := tree.Search(&tree.Kubernetes, command.Cmds, command.IsLastWorkSpace)
	if result == nil {
		// 如果长度大于3个则直接执行命令
		// k get po -n kube-system
		// k get po
		slog.Debug("命令不存在,执行原始命令", "Raw", command.Raw)
		command.Result, err = execCmdString(command.Raw)
		if err != nil {
			command.Result = fmt.Sprintf("%s\n", err.Error())
		}
	}

	if len(result) > 1 {
		// 只展示推荐命令 不执行
		for k, _ := range result {
			// fmt.Printf("%s\n", k)
			command.Result += fmt.Sprintf("%s\n", k)
		}
	} else if len(result) == 1 {
		for _, v := range result {
			if v.IsShell {
				command.Result, err = execCmdString(v.Shell)
				if err != nil {
					command.Result = fmt.Sprintf("%s\n", err.Error())
					continue
				}
			} else {
				// fmt.Printf("%s\n", k)
				fmt.Println(v)
			}
		}
	}

	// 解析一级命令
	command.ProcessDone = true
	// command.Result = "Process 命令解析完成"
	m.next.execute(command)
	slog.Debug("Process 命令解析完成,进入下一步")
}

func (m *Process) setNext(next Cmd) {
	m.next = next
}
