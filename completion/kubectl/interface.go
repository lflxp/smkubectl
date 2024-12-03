package kubectl

type Cmd interface {
	execute(*Command)
	setNext(Cmd)
}

type Department interface {
	execute(*Patient)
	setNext(Department)
}

func Start(command *Command) {
	raw := Raw{}
	process := Process{}
	out := Out{}

	raw.setNext(&process)
	process.setNext(&out)

	raw.execute(command)
}
