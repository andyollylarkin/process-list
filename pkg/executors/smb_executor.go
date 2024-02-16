package executors

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

type SmbExecutor struct {
	exePath string
	user    string
	pass    string
	host    string
	port    int
}

func NewSmbExecutor(smbExePath, user, pass, host string, port int) *SmbExecutor {
	return &SmbExecutor{
		exePath: smbExePath,
		user:    user,
		pass:    pass,
		host:    host,
		port:    port,
	}
}

func (e *SmbExecutor) Exec() ([]byte, error) {
	clientArgs := []string{
		fmt.Sprintf(`--user=%s%%%s`, e.user, e.pass),
		fmt.Sprintf("//%s:%d", e.host, e.port),
		"--interactive=0",
		"--debuglevel=1",
		"powershell.exe get-process|select Id,ProcessName",
	}

	return e.cmdCallInExecutor(e.exePath, clientArgs)
}

func (e *SmbExecutor) cmdCallInExecutor(cmd string, args []string) ([]byte, error) {
	var stdout bytes.Buffer

	var stderr bytes.Buffer

	c := exec.Command(cmd, args...)
	c.Env = os.Environ()

	stdin, err := c.StdinPipe()
	if err != nil {
		return nil, err
	}

	defer stdin.Close()

	c.Stdout = &stdout
	c.Stderr = &stderr

	err = c.Start()
	if err != nil {
		if len(stderr.String()) == 0 && strings.Contains(strings.ToLower(stdout.String()), "error") {
			stderr = stdout
		}

		return stdout.Bytes(), fmt.Errorf("err: %s; stdout: %s; stderr: %s", err, stdout.String(), stderr.String())
	}

	err = c.Wait()
	if err != nil {
		if len(stderr.String()) == 0 && strings.Contains(strings.ToLower(stdout.String()), "error") {
			stderr = stdout
		}

		return stdout.Bytes(), fmt.Errorf("err: %s; stdout: %s; stderr: %s", err, stdout.String(), stderr.String())
	}

	if strings.Contains(strings.ToLower(stdout.String()), "error") {
		stderr = stdout

		return stdout.Bytes(), fmt.Errorf("stdout: %s; stderr: %s. Exit Code: 909", stdout.String(), stderr.String())
	}

	return stdout.Bytes(), err
}
