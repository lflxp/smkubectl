package kubectl

import (
	"bytes"
	"log/slog"
	"os"
	"os/exec"
)

func execCmd(in string) error {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout // 标准输出
	// var errout bytes.Buffer
	cmd.Stderr = os.Stderr // 标准错误
	// cmd.Stderr = &errout // 标准错误
	err := cmd.Run()
	if err != nil {
		// if &errout != nil {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s errout: %s\n", in, err, errout.String())
		// } else {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s\n", in, err)
		// }
		slog.Error(err.Error(), "CMD", in)
		return err
	}
	return nil
}

func execCmdString(in string) (string, error) {
	cmd := exec.Command("sh", "-c", in)
	cmd.Stdin = os.Stdin
	var out bytes.Buffer
	cmd.Stdout = &out // 标准输出
	var errout bytes.Buffer
	// cmd.Stderr = os.Stderr // 标准错误
	cmd.Stderr = &errout // 标准错误
	err := cmd.Run()
	if err != nil {
		// if &errout != nil {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s errout: %s\n", in, err, errout.String())
		// } else {
		// 	log.Fatalf("cmd.Run() failed with %s err: %s\n", in, err)
		// }
		slog.Error(err.Error(), "CMD", in, "errout", errout.String())
		return "", err
	}
	return out.String(), nil
}
