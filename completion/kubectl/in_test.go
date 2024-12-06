package kubectl

import (
	"strings"
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

	t.Run("k d", func(t *testing.T) {
		cmd := &Command{
			Raw: "k d",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k d", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("suiji", func(t *testing.T) {
		x := "dedscd"
		tmp := strings.Replace(x, "d", "1", 1)
		t.Log(tmp)
	})

	// k logs -n argocd argocd-redis-master-0 -f --tail 100
	// 不能带-f查看日志
	t.Run("k logs -n argocd argocd-redis-master-0 --tail 100 ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k logs -n argocd argocd-redis-master-0 --tail 100 ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k logs -n argocd argocd-redis-master-0 --tail 100 ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k edit ingress ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k edit ingress ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k edit ingress ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k get po -n kube-system rke2-coredns-rke2-coredns-7c596479f9-464tt -c ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k get po -n kube-system rke2-coredns-rke2-coredns-7c596479f9-464tt -c ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k get po -n kube-system rke2-coredns-rke2-coredns-7c596479f9-464tt -c ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	// k logs -f -n 1qaz workspacec1941ce79b9d4fde-6c6f4f9446-cv8nd -c
	t.Run("k logs -f -n 1qaz workspacec1941ce79b9d4fde-6c6f4f9446-cv8nd -c ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k logs -f -n 1qaz workspacec1941ce79b9d4fde-6c6f4f9446-cv8nd -c ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k logs -f -n 1qaz workspacec1941ce79b9d4fde-6c6f4f9446-cv8nd -c ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})

	t.Run("k exec -n b01 workspace128ebf5c93b246c0-76c57cb8f7-g7h9z -c ", func(t *testing.T) {
		cmd := &Command{
			Raw: "k exec -n b01 workspace128ebf5c93b246c0-76c57cb8f7-g7h9z -c ",
		}

		process := Process{}
		raw := Raw{}
		raw.setNext(&process)
		out := Out{}
		process.setNext(&out)

		raw.execute(cmd)
		assert.Equal(t, "k exec -n b01 workspace128ebf5c93b246c0-76c57cb8f7-g7h9z -c ", cmd.Raw)

		assert.Equal(t, true, cmd.RawDone)
		assert.NotEmpty(t, cmd.Cmds)
		assert.NotEmpty(t, cmd.Result)
		// assert.Equal(t, 3, len(cmd.Cmds))
	})
}
