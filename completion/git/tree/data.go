package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

// k get po 第三极数据
// k get -n
var KKindValues = map[string]string{
	"origin": "git branch -a",
}

var KArgValues map[string]*c.TreeNode

func KKind() map[string]*c.TreeNode {
	if KArgValues == nil {
		KArgValues = make(map[string]*c.TreeNode)
		for k, v := range KKindValues {
			KArgValues[k] = &c.TreeNode{
				IsShell:  true,
				Shell:    v,
				Children: KArgs(0),
			}
		}
	}
	return KArgValues
}

// k get -n
// k get po -n
var KArgsValues = map[string]string{
	"-d": "git branch",
}

var KAArgsValues map[string]*c.TreeNode

func KArgs(level int) map[string]*c.TreeNode {
	if level == 4 {
		return nil
	}
	level++
	if KAArgsValues == nil {
		KAArgsValues = make(map[string]*c.TreeNode)
		for k, v := range KArgsValues {
			KAArgsValues[k] = &c.TreeNode{
				IsShell:  true,
				Shell:    v,
				Children: KArgs(level),
			}
		}
	}
	return KAArgsValues
}
