package util

import (
	"bytes"
	"fmt"
	"os/exec"
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
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	time.AfterFunc(expire, func() { // 超时后，强制杀死进程
		if cmd.Process != nil && cmd.Process.Pid != 0 {
			pid := fmt.Sprintf("%d", cmd.Process.Pid)
			out = out + fmt.Sprintf("\nexecute timeout: %v seconds.", expire.Seconds())
			cmd.Process.Kill()
			exec.Command("kill", "-9", pid).Run() // 强制杀死进程
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
