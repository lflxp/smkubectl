package completion

import (
	"fmt"
)

type Out struct {
	next Cmd
}

func (r *Out) Execute(command *Command) {
	fmt.Println(command.Result)
}

func (r *Out) SetNext(next Cmd) {
	r.next = next
}
