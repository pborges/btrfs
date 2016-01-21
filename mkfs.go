package btrfs

import (
	"os/exec"
	"regexp"
	"strconv"
)

type MkfsInfo struct {
	Device     string
	Label      string
	NodeSize   int
	LeafSize   int
	SectorSize int
	Size       SizeInBytes
}

func Mkfs(dev string, label string) (info MkfsInfo, err error) {
	var cmd *exec.Cmd
	if UseSudo {
		cmd = exec.Command("sudo", "mkfs.btrfs", dev, "-f", "--label", label)
	} else {
		cmd = exec.Command("mkfs.btrfs", dev, "-f", "--label", label)
	}

	out, err := cmd.Output()
	if err != nil {
		return
	}
	info.Device = dev
	info.Label = label

	infoRegex, err := regexp.Compile(`nodesize\s(\d+)\sleafsize\s(\d+)\ssectorsize\s(\d+)\ssize\s(\d+.\d+\w{3})`)
	if err != nil {
		return
	}

	matches := infoRegex.FindSubmatch(out)
	if len(matches) != 5 {
		err = unexpectedResult(out)
		return
	}
	info.NodeSize, err = strconv.Atoi(string(matches[1]))
	if err != nil {
		return
	}
	info.LeafSize, err = strconv.Atoi(string(matches[2]))
	if err != nil {
		return
	}
	info.SectorSize, err = strconv.Atoi(string(matches[3]))
	if err != nil {
		return
	}

	info.Size, err = parseSize(matches[4])
	if err != nil {
		return
	}
	return
}