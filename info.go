package btrfs

import (
	"regexp"
	"bytes"
	"strconv"
	"io"
	"fmt"
)

type Array struct {
	Label   string
	UUID    string
	Devices []Device
}

func (s SizeInBytes) String() string {
	if s > TeraByteMultiplicant {
		return fmt.Sprintf("%4.2fTiB", float64(s / TeraByteMultiplicant))
	} else if s > GigabyteMultiplicant {
		return fmt.Sprintf("%4.2fGiB", float64(s / GigabyteMultiplicant))
	} else if s > MegabyteMultiplicant {
		return fmt.Sprintf("%4.2fMiB", float64(s / MegabyteMultiplicant))
	} else if s > KilobyteMultiplicant {
		return fmt.Sprintf("%4.2fKiB", float64(s / KilobyteMultiplicant))
	}
	return fmt.Sprintf("%4.2fB", float64(s))
}

type Device struct {
	Id     int
	Size   SizeInBytes
	Used   SizeInBytes
	Free   SizeInBytes
	Device string
}

func Info() (arrays []Array, err error) {
	arrays = make([]Array, 0)
	out, err := getCommand("btrfs", "filesystem", "show").Output()
	if err != nil {
		return
	}
	r := bytes.NewBuffer(out)

	labelRegex, err := regexp.Compile(`Label:\s+(?:')?(\w+)(?:')?\s+uuid:\s(\w{8}-\w{4}-\w{4}-\w{4}-\w{12})`)
	if err != nil {
		return
	}
	deviceRegex, err := regexp.Compile(`devid\s+(\d)\ssize\s(\d+\.\d+\w+)\sused\s+(\d+\.\d+\w+)\s+path\s+([\w/]+)`)
	if err != nil {
		return
	}

	var line []byte
	var a *Array
	for ; err == nil; line, err = r.ReadBytes(0x0a) {
		if arrayMatches := labelRegex.FindSubmatch(line); len(arrayMatches) == 3 {
			if a != nil {
				arrays = append(arrays, *a)
			}
			a = new(Array)
			a.Label = string(arrayMatches[1])
			a.UUID = string(arrayMatches[2])
			a.Devices = make([]Device, 0)
		}else if deviceMatches := deviceRegex.FindSubmatch(line); len(deviceMatches) == 5 {
			d := Device{}
			d.Id, err = strconv.Atoi(string(deviceMatches[1]))
			if err != nil {
				return
			}

			d.Size, err = parseSize(deviceMatches[2])
			if err != nil {
				return
			}
			d.Used, err = parseSize(deviceMatches[3])
			if err != nil {
				return
			}

			d.Free = d.Size - d.Used
			d.Device = string(deviceMatches[4])
			a.Devices = append(a.Devices, d)
		}
	}
	if err == io.EOF {
		err = nil
	}
	if a != nil {
		arrays = append(arrays, *a)
	} else {
		err = unexpectedResult(out)
	}
	return
}