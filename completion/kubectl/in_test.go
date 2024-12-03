package kubectl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_raw_execute(t *testing.T) {
	t.Run("", func(t *testing.T) {
		cmd := &Command{
			Raw: "",
		}

		raw := Raw{}

		raw.execute(cmd)
		assert.Equal(t, "", cmd.Raw)

		assert.Equal(t, false, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
		assert.Equal(t, []string{}, cmd.Cmds)

		process := Process{}
		raw.setNext(&process)

		out := Out{}
		process.setNext(&out)

		out.execute(cmd)
	})

	t.Run("k", func(t *testing.T) {
		cmd := &Command{
			Raw: "k",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)

		assert.Equal(t, "k", cmd.Raw)
		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
		assert.Equal(t, []string{"k"}, cmd.Cmds)
	})

	t.Run("k ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, true, cmd.IsLastWorkSpace)
		assert.Equal(t, 1, len(cmd.Cmds))
	})

	t.Run("k get -n ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k get -n ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k get -n ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, true, cmd.IsLastWorkSpace)
		assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k get po -", func(t *testing.T) {
		cmd := &Command{
			Raw: "k get po -",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k get po -", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k get po", func(t *testing.T) {
		cmd := &Command{
			Raw: "k get po",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k get po", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k get po ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k get po ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k get po ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})
}
