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
	// 需要加上这一行，以便在windows下能够正常执行，而不会不停弹窗
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	time.AfterFunc(expire, func() { // 超时后，强制杀死进程
		if cmd.Process != nil && cmd.Process.Pid != 0 {
			out = out + fmt.Sprintf("\nexecute timeout: %v seconds.", expire.Seconds())
			cmd.Process.Kill()
		}
	})

	err = cmd.Run()
	if err != nil {
		err = fmt.Errorf("%v\n%v\n%v", err.Error(), stderr.String(), out)
		return
	}
	out = stdout.String()
	return
}
