package completion

type Cmd interface {
	Execute(*Command)
	SetNext(Cmd)
}

type Command struct {
	IsLastWorkSpace bool     // 记录命令最后一个字符是否是空格
	Raw             string   // 原始命令
	Cmds            []string // 命令
	RawDone         bool     // 原始命令是否已经完成
	ProcessDone     bool     // 方法是否已经完成
	Result          string   // 结果
}
