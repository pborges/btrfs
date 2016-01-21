package main

import (
	"github.com/pborges/btrfs"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	btrfs.UseSudo = true

	log.Println("MKFS --------------------------------------------------")
	f, err := btrfs.Mkfs("/dev/sdb", "disk0")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(f)

	f, err = btrfs.Mkfs("/dev/sdc", "disk1")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(f)

	info()

	log.Println("MOUNT--------------------------------------------------")
	err = btrfs.Mount("/dev/sdb", "tmp")
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("ADD  --------------------------------------------------")
	err = btrfs.Add("/dev/sdc", "tmp")
	if err != nil {
		log.Fatalln(err)
	}

	info()

	log.Println("DELETE-------------------------------------------------")
	err = btrfs.Delete("/dev/sdc", "tmp")
	if err != nil {
		log.Fatalln(err)
	}

	info()

	log.Println("UMOUNT-------------------------------------------------")
	err = btrfs.Umount("tmp")
	if err != nil {
		log.Fatalln(err)
	}

}

func info() {
	x, err := btrfs.Info()
	if err != nil {
		log.Fatalln(err)
	}
	for _, a := range x {
		log.Printf("Label : %s\n", a.Label)
		log.Printf("UUID  : %s\n", a.UUID)
		for _, d := range a.Devices {
			log.Printf("Id    : %d\n", d.Id)
			log.Printf("\tUsed  : %s\n", d.Used)
			log.Printf("\tFree  : %s\n", d.Free)
			log.Printf("\tDevice: %s\n", d.Device)
		}
	}
}