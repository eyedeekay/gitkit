package gitkit

import (
	psutilproc "github.com/shirou/gopsutil/process"
	"io"
	"os"
	"os/exec"
)

func cleanUpProcessGroup(cmd *exec.Cmd) {
	if cmd == nil {
		return
	}

	process := cmd.Process
	if process != nil && process.Pid > 0 {
		p, _ := psutilproc.NewProcess(int32(process.Pid)) // Specify process id of parent
		r, _ := p.Children()
		for _, v := range r {
			v.Kill()
		}

		p.Kill()
	}

	go cmd.Wait()
}

func gitCommand(name string, args ...string) (*exec.Cmd, io.Reader) {
	cmd := exec.Command(name, args...)
	//	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true}
	cmd.Env = os.Environ()

	r, _ := cmd.StdoutPipe()
	cmd.Stderr = cmd.Stdout

	return cmd, r
}
