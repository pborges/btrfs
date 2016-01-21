package btrfs

import (
	"strconv"
	"regexp"
)

type SizeInBytes int64

func parseSize(rawSize []byte) (size SizeInBytes, err error) {
	sizeRegex, err := regexp.Compile(`(\d+.\d+)(\w+)`)
	if err != nil {
		return
	}
	matches := sizeRegex.FindSubmatch(rawSize)

	s, err := strconv.ParseFloat(string(matches[1]), 64)
	if err != nil {
		return
	}
	size = SizeInBytes(s * float64(sizeMap[string(matches[2])]))
	return
}
