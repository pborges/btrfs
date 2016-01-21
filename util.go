package btrfs

import "os/exec"

func getCommand(args... string) (cmd *exec.Cmd) {
	if len(args) <= 0 {
		return
	}

	if UseSudo {
		cmd = exec.Command("sudo", args...)
	} else {
		cmd = exec.Command(args[0], args[1:]...)
	}
	return
}