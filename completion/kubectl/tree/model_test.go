package tree

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func Test_Search(t *testing.T) {
	t.Run("k", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k"}, false)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 2)
		assert.Equal(t, rs["k"].IsShell, true)
	})

	t.Run("k get", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get"}, false)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 1)
		assert.Equal(t, rs["get"].IsShell, true)
	})

	t.Run("k get ", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get"}, true)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 1)
		assert.Equal(t, rs["get"].IsShell, true)
	})

	t.Run("k a", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "a"}, false)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 2)
		assert.Equal(t, rs["api-resources"].IsShell, false)
	})

	t.Run("k get -n ", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get", "-n"}, true)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 1)
		assert.Equal(t, rs["-n"].IsShell, true)
	})

	t.Run("k get po -n ", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get", "po", "-n"}, true)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 1)
		assert.Equal(t, rs["-n"].IsShell, true)
	})

	t.Run("k get cm -n ", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get", "cm", "-n"}, true)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 1)
		assert.Equal(t, rs["-n"].IsShell, true)
	})

	t.Run("k get cm -", func(t *testing.T) {
		rs := Search(&Kubernetes, []string{"k", "get", "cm", "-"}, false)
		if rs == nil {
			t.Fatal("rs is nil")
		}

		assert.Equal(t, len(rs), 2)
		assert.Equal(t, rs["-n"].IsShell, true)
	})
}
