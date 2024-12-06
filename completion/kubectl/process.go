package kubectl

import (
	"fmt"
	"strings"

	"log/slog"

	"github.com/lflxp/smkubectl/completion/kubectl/tree"
)

// https://refactoringguru.cn/design-patterns/chain-of-responsibility/go/example#example-0
const CC = `$HERE -o=jsonpath='{range .spec.containers[*]}{.name}{"\n"}{end}'`

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
		// 原始命令解析
		command.Result, err = execCmdString(command.Raw)
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
				// 如果最后的命令是-c则执行
				if command.Cmds[len(command.Cmds)-1] == "-c" {
					var cmd string
					cmd = strings.ReplaceAll(CC, "$HERE", strings.Join(command.Cmds[:len(command.Cmds)-1], " "))
					switch command.Cmds[1] {
					case "logs":
						if command.Cmds[len(command.Cmds)-1] == "-c" {
							// 判断-n后面是否有命名空间
							cmd = command.Raw
							for index, v := range command.Cmds {
								if v == "-n" {
									if len(command.Cmds) >= index+2 {
										ns := command.Cmds[index+1]
										pod := command.Cmds[index+2]
										cmd = fmt.Sprintf("kubectl get po -n %s %s", ns, pod)
										cmd = strings.ReplaceAll(CC, "$HERE", cmd)
									}
								}
							}
						}
					case "exec":
						if command.Cmds[len(command.Cmds)-1] == "-c" {
							// 判断-n后面是否有命名空间
							cmd = command.Raw
							for index, v := range command.Cmds {
								if v == "-n" {
									if len(command.Cmds) >= index+2 {
										ns := command.Cmds[index+1]
										pod := command.Cmds[index+2]
										cmd = fmt.Sprintf("kubectl get po -n %s %s", ns, pod)
										cmd = strings.ReplaceAll(CC, "$HERE", cmd)
									}
								}
							}
						}
					}

					command.Result, err = execCmdString(cmd)
					command.Result = fmt.Sprintf("执行命令: %s\n%s\n", cmd, command.Result)
				} else {
					command.Result, err = execCmdString(v.Shell)
				}

				if err != nil {
					command.Result = fmt.Sprintf("Error\n%s\n", err.Error())
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
