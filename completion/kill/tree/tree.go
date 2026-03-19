package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

var Kill = c.TreeNode{
	Value:  "kill",
	Kind:   c.KindCommand,
	Common: "Kill process command",
	Children: map[string]*c.TreeNode{
		"kill": {
			IsShell:  false,
			Children: KillCommand,
		},
	},
}

var KillCommand = map[string]*c.TreeNode{
	"-9": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
	"-KILL": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
	"-15": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
	"-TERM": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
	"-1": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
	"-HUP": {
		IsShell: true,
		Shell:   "ps -eo pid,uid,ppid,pcpu,stime,tty,time,args",
	},
}
