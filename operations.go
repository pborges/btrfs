package btrfs

func Add(dev string, mount string) (err error) {
	_, err = getCommand("btrfs", "device", "add", "-f", dev, mount).Output()
	return
}

func Delete(dev string, mount string) (err error) {
	_, err = getCommand("btrfs", "device", "delete", dev, mount).Output()
	return
}
