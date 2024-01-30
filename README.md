# 介绍

![](asset/pod.png)

smkubectl = kubectl + fzf + zsh-completion 是一个用于在 Kubernetes 环境中执行命令的工具。Kubernetes 是一个用于管理容器化应用程序的操作系统，它运行在分布式集群上。kubectl 是用于与 Kubernetes 集群交互的命令行工具，而 fzf 是一个用于在文本文件中执行交互式搜索的工具，而 zsh-completion 是用于自动完成命令的补全工具。

这个工具组合的设计是为了提高用户在 Kubernetes 环境中的命令执行效率。kubectl 提供了丰富的命令集，用于管理 Kubernetes 集群上的各种资源。但是，这些命令通常很长，在终端中一次输入所有命令可能会很麻烦。fzf 提供了交互式搜索功能，允许用户通过键盘快捷键或者模糊搜索来选择所需的命令。最后，zsh-completion 可以根据用户输入的命令前缀自动完成命令，减少用户输入的工作量。

# 依赖

要使用这个工具组合，你需要安装并配置 kubectl、fzf 和 zsh。以下是安装和配置的步骤：

* 安装 kubectl
* 安装 fzf
* 安装 zsh

通过使用 kubectl + fzf + zsh-completion 工具组合，你可以提高在 Kubernetes 环境中执行命令的效率，使你的工作更轻松。

# 安装 

1. install

```
go install github.com/lflxp/smkubectl@latest
```

2. zsh开启自动补全

```
autoload -U compinit && compinit -u
source <(smkubectl completion zsh)
ln -s `which kubectl` /usr/local/bin/kubectl
```

# 使用

## 快捷键

> ~

## 操作

* k + ~
* k g + ~
* k get + ~
* k get po+~ (没有空格)
* k get po + ~ (有空格)
* k edit po -n 
* k get po -n namespace pod -c + ~
* k logs -f + ~