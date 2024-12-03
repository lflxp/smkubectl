package cmd

import "testing"

func Test_ParseCmd(t *testing.T) {
	ParseCmd("kubectl get pod")
}
