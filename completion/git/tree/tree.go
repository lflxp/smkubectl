package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

var Gits = c.TreeNode{
	Value:  "g",
	Kind:   c.KindCommand,
	Common: "Git 命令行工具",
	Children: map[string]*c.TreeNode{
		"g": {
			IsShell:  true,
			Shell:    "git",
			Children: GCommand,
		},
		"git": {
			IsShell:  true,
			Shell:    "git",
			Children: GCommand,
		},
		"gs": {
			IsShell:  true,
			Shell:    "git status",
			Children: GCommand,
		},
		"gl": {
			IsShell:  true,
			Shell:    "git l2",
			Children: GCommand,
		},
	},
}

var GCommand = map[string]*c.TreeNode{
	"status": {
		IsShell: true,
		Shell:   "git status",
	},
	"log": {
		IsShell: true,
		Shell:   `git l2`,
	},
	"branch": {
		IsShell:  true,
		Shell:    `git branch -vv`,
		Children: KArgs(0),
	},
	"checkout": {
		IsShell: true,
		Shell:   `git branch -vv`,
	},
	"switch": {
		IsShell: true,
		Shell:   `git branch -vv`,
	},
	"pull": {
		IsShell:  true,
		Shell:    `git remote`,
		Children: KKind(),
	},
	"push": {
		IsShell:  true,
		Shell:    `git remote`,
		Children: KKind(),
	},
	"add": {
		IsShell:  true,
		Shell:    `git status`,
		Children: KKind(),
	},
	"reset": {
		IsShell: true,
		Shell:   `git reset --hard HEAD`,
	},
	"merge": {
		IsShell: true,
		Shell:   `git branch -a`,
	},
	"remote": {
		IsShell: false,
		Children: map[string]*c.TreeNode{
			"add": {
				IsShell: true,
				Shell:   `git remote`,
			},
			"remove": {
				IsShell: true,
				Shell:   `git remote`,
			},
		},
	},
}
