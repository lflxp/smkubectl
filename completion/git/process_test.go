package git

import (
	"testing"

	c "github.com/lflxp/smkubectl/completion"
	"github.com/stretchr/testify/assert"
)

func Test_raw_Execute(t *testing.T) {
	t.Run("", func(t *testing.T) {
		cmd := &c.Command{
			Raw: "",
		}

		raw := c.Raw{}

		raw.Execute(cmd)
		assert.Equal(t, "", cmd.Raw)

		assert.Equal(t, false, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
		assert.Equal(t, []string{}, cmd.Cmds)

		process := Process{}
		raw.SetNext(&process)

		out := c.Out{}
		process.SetNext(&out)

		out.Execute(cmd)
	})

	t.Run("g", func(t *testing.T) {
		cmd := &c.Command{
			Raw: "g",
		}

		process := Process{}
		raw := c.Raw{}
		raw.SetNext(&process)
		out := c.Out{}
		process.SetNext(&out)

		raw.Execute(cmd)

		assert.Equal(t, "g", cmd.Raw)
		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
		assert.Equal(t, []string{"g"}, cmd.Cmds)
	})

	t.Run("g ", func(t *testing.T) {
		cmd := &c.Command{
			Raw: "g ",
		}

		process := Process{}
		raw := c.Raw{}
		raw.SetNext(&process)
		out := c.Out{}
		process.SetNext(&out)

		raw.Execute(cmd)

		assert.Equal(t, "g ", cmd.Raw)
		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
		assert.Equal(t, []string{"g"}, cmd.Cmds)
	})

	t.Run("g branch -d ", func(t *testing.T) {
		cmd := &c.Command{
			Raw: "g branch -d ",
		}

		process := Process{}
		raw := c.Raw{}
		raw.SetNext(&process)
		out := c.Out{}
		process.SetNext(&out)

		raw.Execute(cmd)

		assert.Equal(t, "g branch -d ", cmd.Raw)
		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
	})

	t.Run("g remote ", func(t *testing.T) {
		cmd := &c.Command{
			Raw: "g remote ",
		}

		process := Process{}
		raw := c.Raw{}
		raw.SetNext(&process)
		out := c.Out{}
		process.SetNext(&out)

		raw.Execute(cmd)

		assert.Equal(t, "g remote ", cmd.Raw)
		assert.Equal(t, true, cmd.RawDone)
		assert.Equal(t, false, cmd.IsLastWorkSpace)
	})
}
