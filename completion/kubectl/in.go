package kubectl

import (
	"fmt"
	"log/slog"
	"strings"
)

type Raw struct {
	next Cmd
}

func (r *Raw) execute(command *Command) {
	if command.RawDone {
		r.next.execute(command)
		return
	}

	// 解析命令
	tmp := strings.Split(command.Raw, " ")
	result := []string{}
	for index, v := range tmp {
		if v != "" && v != " " {
			result = append(result, v)
		}
		if index == len(tmp)-1 && v == "" {
			slog.Debug("最后一个字符是空格")
			command.IsLastWorkSpace = true
		}
	}

	command.Cmds = result
	r.show(command)

}

func (r *Raw) setNext(next Cmd) {
	r.next = next
}

// 提示命令
// 显示结果
func (r *Raw) show(command *Command) {
	if command.Raw == "" {
		fmt.Println("请输入关键字: git | go | kubectl | k | showme")
		return
	}

	command.RawDone = true

	// 判断是走到method方法还是提示
	if r.next != nil {
		slog.Debug("In 命令解析完成,进入下一步")
		r.next.execute(command)
	}
}
