package bytesex

import "math"

// ToKiB to KiB
func ToKiB(b uint64) float64 {
	return math.Round(float64(b)/1024*100) / 100
}

// ToMiB to MiB
func ToMiB(b uint64) float64 {
	// 1024x1024
	return math.Round(float64(b)/1048576*100) / 100
}

// ToGiB to GiB
func ToGiB(b uint64) float64 {
	// 1024x1024x1024
	return math.Round(float64(b)/1073741824*100) / 100
}

// ToTiB to TiB
func ToTiB(b uint64) float64 {
	// 1024x1024x1024x1024
	return math.Round(float64(b)/1099511627776*100) / 100
}
