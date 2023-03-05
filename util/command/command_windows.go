package command

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"syscall"
	"time"
)

// ExecCommand 执行cmd命令操作
func ExecCommand(name string, args []string, timeout ...time.Duration) (out string, err error) {
	var (
		stderr, stdout bytes.Buffer
		expire         = 30 * time.Minute
		errs           []string
	)

	if len(timeout) > 0 {
		expire = timeout[0]
	}

	cmd := exec.Command(name, args...)
	// 需要加上这一行，以便在windows下能够正常执行，而不会不停弹窗
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	err = cmd.Start()
	if err != nil {
		err = fmt.Errorf("%s\n%s", err.Error(), stderr.String())
		return
	}

	time.AfterFunc(expire, func() {
		if cmd.Process != nil && cmd.Process.Pid != 0 {
			errs = append(errs, fmt.Sprintf("execute timeout: %d min.", int(expire.Minutes())))
			if proc, errProc := os.FindProcess(cmd.Process.Pid); errProc == nil {
				proc.Kill()
				proc.Release()
			}
			cmd.Process.Kill()
		}
	})

	err = cmd.Wait()
	if err != nil {
		errs = append(errs, err.Error(), stderr.String())
	}
	out = stdout.String()
	if len(errs) > 0 {
		errs = append(errs, out)
		err = errors.New(strings.Join(errs, "\n\r"))
	}
	return
}
