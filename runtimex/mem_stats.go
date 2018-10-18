package runtimex

import (
	"fmt"
	"runtime"

	"github.com/chouandy/goex/bytesex"
)

// PrintMemStats print mem stats
func PrintMemStats() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Println(fmt.Sprintf(
		"Alloc: %.2f MiB\tTotalAlloc: %.2f MiB\tSys: %.2f MiB\tNumGC: %v\tNumForcedGC: %v",
		bytesex.ToMiB(m.Alloc),
		bytesex.ToMiB(m.TotalAlloc),
		bytesex.ToMiB(m.Sys),
		m.NumGC, m.NumForcedGC,
	))
}
