package kubectl

import (
	"fmt"
)

type Out struct {
	next Cmd
}

func (r *Out) execute(command *Command) {
	fmt.Println(command.Result)
}

func (r *Out) setNext(next Cmd) {
	r.next = next
}
