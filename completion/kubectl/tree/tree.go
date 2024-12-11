package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

var Kubernetes = c.TreeNode{
	Value:  "k",
	Kind:   c.KindCommand,
	Common: "Kubernetes 命令行工具",
	Children: map[string]*c.TreeNode{
		"k": {
			IsShell:  true,
			Shell:    "kubectl",
			Children: KCommand,
		},
		"kubectl": {
			IsShell:  true,
			Shell:    "kubectl",
			Children: KCommand,
		},
	},
}

var KCommand = map[string]*c.TreeNode{
	"get": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"scale": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"edit": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"patch": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"delete": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"label": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"explain": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KArgs(0),
	},
	"api-resources": {
		IsShell: true,
		Shell:   "kubectl api-resources",
	},
	"api-versions": {
		IsShell: true,
		Shell:   "kubectl api-versions",
	},
	"logs": {
		IsShell:  true,
		Shell:    "kubectl get po -A",
		Children: KArgs(0),
	},
	"describe": {
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"exec": {
		IsShell:  true,
		Shell:    "kubectl get po -A",
		Children: KArgs(0),
	},
	"-n": {
		IsShell: true,
		Shell:   "kubectl get ns",
	},
}
