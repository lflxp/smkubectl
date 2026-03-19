package tree

import (
	c "github.com/lflxp/smkubectl/completion"
)

var Docker = c.TreeNode{
	Value:  "docker",
	Kind:   c.KindCommand,
	Common: "Docker 命令行工具",
	Children: map[string]*c.TreeNode{
		"docker": {
			IsShell:  false,
			Shell:    "docker",
			Children: DockerCommand,
		},
		"d": {
			IsShell:  false,
			Shell:    "docker",
			Children: DockerCommand,
		},
	},
}

var DockerCommand = map[string]*c.TreeNode{
	"exec": {
		IsShell:  true,
		Shell:    "docker ps",
		Children: DockerExecArgs,
	},
	"logs": {
		IsShell:  true,
		Shell:    "docker ps --format '{{.Names}}'",
		Children: DockerLogsArgs,
	},
	"inspect": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: DockerInspectArgs,
	},
	"run": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: DockerRunArgs,
	},
	"ps": {
		IsShell: true,
		Shell:   "docker ps --format '{{.Names}}'",
	},
	"images": {
		IsShell: true,
		Shell:   "docker images --format '{{.Repository}}:{{.Tag}}'",
	},
	"stop": {
		IsShell: true,
		Shell:   "docker ps --format '{{.Names}}'",
	},
	"start": {
		IsShell: true,
		Shell:   "docker ps -a --format '{{.Names}}'",
	},
	"rm": {
		IsShell: true,
		Shell:   "docker ps -a --format '{{.Names}}'",
	},
	"rmi": {
		IsShell: true,
		Shell:   "docker images --format '{{.Repository}}:{{.Tag}}'",
	},
	"pull": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}' | sort -u",
		Children: nil,
	},
	"push": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"build": {
		IsShell: true,
		Shell:   "docker images --format '{{.Repository}}:{{.Tag}}'",
	},
	"tag": {
		IsShell: true,
		Shell:   "docker images --format '{{.Repository}}:{{.Tag}}'",
	},
	"network": {
		IsShell: true,
		Shell:   "docker network ls --format '{{.Name}}'",
	},
	"volume": {
		IsShell: true,
		Shell:   "docker volume ls --format '{{.Name}}'",
	},
}

var DockerExecArgs = map[string]*c.TreeNode{
	"-it": {
		IsShell:  true,
		Shell:    "docker ps",
		Children: nil,
	},
	"-i": {
		IsShell:  true,
		Shell:    "docker ps --format '{{.Names}}'",
		Children: nil,
	},
	"-t": {
		IsShell:  true,
		Shell:    "docker ps --format '{{.Names}}'",
		Children: nil,
	},
}

var DockerLogsArgs = map[string]*c.TreeNode{
	"-f": {
		IsShell:  true,
		Shell:    "docker ps",
		Children: nil,
	},
	"--follow": {
		IsShell:  true,
		Shell:    "docker ps",
		Children: nil,
	},
}

var DockerInspectArgs = map[string]*c.TreeNode{
	"container": {
		IsShell:  true,
		Shell:    "docker ps --format '{{.Names}}'",
		Children: nil,
	},
	"image": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"network": {
		IsShell:  true,
		Shell:    "docker network ls --format '{{.Name}}'",
		Children: nil,
	},
	"volume": {
		IsShell:  true,
		Shell:    "docker volume ls --format '{{.Name}}'",
		Children: nil,
	},
}

var DockerRunArgs = map[string]*c.TreeNode{
	"-it": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"-d": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"--rm": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"-p": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"-v": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"-e": {
		IsShell:  true,
		Shell:    "docker images --format '{{.Repository}}:{{.Tag}}'",
		Children: nil,
	},
	"--network": {
		IsShell:  true,
		Shell:    "docker network ls --format '{{.Name}}'",
		Children: nil,
	},
}
