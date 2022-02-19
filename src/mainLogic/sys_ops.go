package mainLogic

import (
	"ControllerGo/src/osSpecific"
	"fmt"
	"os"
	"os/exec"
)

var proc *os.Process

func CheckEventServer() {
	if proc == nil {
		panic(osSpecific.EventServerNotRunningMsg)
	}
	osSpecific.CheckProcess(proc.Pid)
}

func StartProcess(execPath string) {
	cmd := exec.Command(execPath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Start()
	CheckErr(err)
	proc = cmd.Process
	pid := proc.Pid
	CheckEventServer()
	osSpecific.SetHighPriority(pid)
	fmt.Printf("Event service started. PID: %d\n", pid)
}

func StartEventServer() {
	StartProcess(EventServerExecPath)
}

func KillEventServer() {
	if proc != nil {
		proc.Kill()
	}
}

func SetSelfPriority() {
	pid := os.Getppid()
	osSpecific.SetHighPriority(pid)
}