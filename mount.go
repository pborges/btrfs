package btrfs

import (
	"os/exec"
	"os"
)

func Mount(dev string, mount string) (err error) {
	err = os.Mkdir(mount, 0755)
	if _, ok := err.(*os.PathError); err != nil && !ok {
		return
	}
	var cmd *exec.Cmd
	if UseSudo {
		cmd = exec.Command("sudo", "mount", dev, mount)
	} else {
		cmd = exec.Command("mount", dev, mount)
	}

	_, err = cmd.Output()
	return
}

func Umount(mount string) (err error) {
	var cmd *exec.Cmd
	if UseSudo {
		cmd = exec.Command("sudo", "umount", mount)
	} else {
		cmd = exec.Command("umount", mount)
	}

	_, err = cmd.Output()
	if err != nil {
		return
	}
	err = os.Remove(mount)
	return
}