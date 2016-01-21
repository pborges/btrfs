package btrfs

import "strconv"

type SizeInBytes int64

func parseSize(rawSize []byte) (size SizeInBytes, err error) {
	n := len(rawSize) - 3
	unit := string(rawSize[n:])
	s, err := strconv.ParseFloat(string(rawSize[:n]), 64)
	if err != nil {
		return
	}
	size = SizeInBytes(s * float64(sizeMap[unit]))
	return
}
