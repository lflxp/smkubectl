package tree

import "strings"

type Kind string

const (
	KindCommand Kind = "command"
	KindArg     Kind = "arg"
)

type TreeNode struct {
	Value    string   // 命令
	Kind     Kind     // 类型：命令、参数
	Common   string   // 备注
	IsShell  bool     // 是否执行kubectl获取还是直接cmd提示
	Shell    string   // 获取提示的命令
	Cmds     []string // 命令
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

// 按数组顺序搜索
// 判断最后一个字符是否为空格
// 如果最后一个字符不为空格，搜索到最后一个关键字支持模糊搜索
// 递归查询
func Search(tree *TreeNode, data []string, isLastSpace bool) map[string]*TreeNode {
	if len(data) == 0 {
		return nil
	}

	// 支持模糊匹配
	if len(data) == 1 {
		if isLastSpace {
			// 查询空格前一个命令的精确匹配
			if node, ok := tree.Children[data[0]]; ok {
				return map[string]*TreeNode{data[0]: node}
			}
			return nil
		} else {
			// 进行命令的模糊搜索
			rs := make(map[string]*TreeNode)
			for k, v := range tree.Children {
				if strings.HasPrefix(k, data[0]) {
					rs[k] = v
				}
			}
			return rs
		}
	}

	// 递归查询
	if node, ok := tree.Children[data[0]]; ok {
		return Search(node, data[1:], isLastSpace)
	}

	return nil
}
