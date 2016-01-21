package main

import (
	"github.com/pborges/btrfs"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	btrfs.UseSudo = true

	f, err := btrfs.Mkfs("/dev/sdb", "disk0")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Mkfs:", f)

	f, err = btrfs.Mkfs("/dev/sdc", "disk1")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Mkfs:", f)

	x, err := btrfs.Info()
	if err != nil {
		log.Fatalln(err)
	}
	for _, a := range x {
		log.Printf("Label: %s\n", a.Label)
		log.Printf("UUID: %s\n", a.UUID)
		for _, d := range a.Devices {
			log.Printf("Device %d\n", d.Id)
			log.Printf("Used: %s\n", d.Used)
			log.Printf("Free: %s\n", d.Free)
			log.Printf("Path: %s\n", d.Path)

		}
	}
}
