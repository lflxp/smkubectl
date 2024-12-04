package tree

var Kubernetes = TreeNode{
	Value:  "k",
	Kind:   KindCommand,
	Common: "Kubernetes 命令行工具",
	Children: map[string]*TreeNode{
		"k": &TreeNode{
			IsShell:  true,
			Shell:    "kubectl",
			Children: KCommand,
		},
		"kubectl": &TreeNode{
			IsShell:  true,
			Shell:    "kubectl",
			Children: KCommand,
		},
	},
}

var KCommand = map[string]*TreeNode{
	"get": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"edit": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"patch": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"delete": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"label": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl api-resources",
		Children: KKind(),
	},
	"api-resources": &TreeNode{
		IsShell: false,
	},
	"api-versions": &TreeNode{
		IsShell: false,
	},
	"logs": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl get po -A",
		Children: KArgs(0),
	},
	"describe": &TreeNode{
		IsShell:  true,
		Shell:    "kubectl get po -A",
		Children: KKind(),
	},
	"-n": &TreeNode{
		IsShell: true,
		Shell:   "kubectl get ns",
	},
}
