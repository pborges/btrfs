package btrfs

var UseSudo bool = false
var sizeMap map[string]int64

const (
	TeraByteMultiplicant = 1000000000000
	GigabyteMultiplicant = 1000000000
	MegabyteMultiplicant = 1000000
	KilobyteMultiplicant = 1000
)

func init() {
	sizeMap = make(map[string]int64)
	sizeMap["TiB"] = TeraByteMultiplicant
	sizeMap["GiB"] = GigabyteMultiplicant
	sizeMap["MiB"] = MegabyteMultiplicant
	sizeMap["KiB"] = KilobyteMultiplicant
}