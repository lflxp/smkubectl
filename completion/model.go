package completion

import (
	"strings"
)

type Kind string

const (
	KindCommand Kind = "command"
	KindArg     Kind = "arg"
)

type TreeNode struct {
	Value    string
	Kind     Kind
	Common   string
	IsShell  bool // ### IsShell 的作用 IsShell 值 行为 false 不执行命令，只显示 Children 中的子命令列表 true 执行 Shell 字段中的命令，返回命令输出作为补全选项
	Shell    string
	Cmds     []string
	IsHeader bool
	Header   string
	Children map[string]*TreeNode
}

func (t *TreeNode) AddChild(node *TreeNode) {
	if t.Children == nil {
		t.Children = make(map[string]*TreeNode)
	}
	t.Children[node.Value] = node
}

func (t *TreeNode) GetChild(value string) *TreeNode {
	if t.Children == nil {
		return nil
	}
	return t.Children[value]
}

func Search(tree *TreeNode, data []string, isLastSpace bool) map[string]*TreeNode {
	if len(data) == 0 {
		if tree.IsShell && tree.Shell != "" {
			return map[string]*TreeNode{"": tree}
		}
		return nil
	}

	if len(data) == 1 {
		if isLastSpace {
			if node, ok := tree.Children[data[0]]; ok {
				return map[string]*TreeNode{data[0]: node}
			}
			return nil
		} else {
			rs := make(map[string]*TreeNode)
			for k, v := range tree.Children {
				if strings.HasPrefix(k, data[0]) {
					rs[k] = v
				}
			}
			return rs
		}
	}

	if node, ok := tree.Children[data[0]]; ok {
		return Search(node, data[1:], isLastSpace)
	} else {
		if len(data) > 1 {
			return Search(tree, data[2:], isLastSpace)
		}
	}

	return nil
}
