package tree

import (
	"testing"

	c "github.com/lflxp/smkubectl/completion"
	"github.com/magiconair/properties/assert"
)

func Test_Search(t *testing.T) {
	t.Run("g", func(t *testing.T) {
		rs := c.Search(&Gits, []string{"g"}, false)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 2)
		assert.Equal(t, rs["g"].IsShell, true)
	})
}
