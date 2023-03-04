package util

import (
	"bytes"
	"fmt"
	"os/exec"
	"syscall"
	"time"
)

// ExecCommand 执行cmd命令操作
func ExecCommand(name string, args []string, timeout ...time.Duration) (out string, err error) {
	var (
		stderr, stdout bytes.Buffer
		expire         = 30 * time.Minute
	)

	if len(timeout) > 0 {
		expire = timeout[0]
	}

	cmd := exec.Command(name, args...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	isDone := make(chan bool, 1)
	go func() {
		err = cmd.Run()
		if err != nil {
			err = fmt.Errorf("%s\n%s", err.Error(), stderr.String())
		}
		isDone <- true
	}()

	select {
	case <-isDone:
	case <-time.After(expire):
		if cmd.Process != nil && cmd.Process.Pid != 0 {
			err = fmt.Errorf("execute timeout: %v seconds.", expire.Seconds())
			syscall.Kill(-cmd.Process.Pid, syscall.SIGKILL)
			cmd.Process.Kill()
		}
	}
	out = stdout.String()
	return
}
