package util

import (
	"fmt"
	"runtime"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func PrintMemUsage(num int) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	rl.DrawText(fmt.Sprintf("Allocated objects = %v", num), 500, 100, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("\tTotalAlloc on Heap = %v MiB", bToMb(m.TotalAlloc)), 500, 200, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("\tSys = %v MiB", bToMb(m.Sys)), 500, 300, 20, rl.Red)
	rl.DrawText(fmt.Sprintf("\tNumGC = %v\n", m.NumGC), 500, 400, 20, rl.Red)
}

func PrintMemUsageConsole() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	fmt.Printf("Alloc = %v MiB", bToMb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", bToMb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", bToMb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func bToMb(b uint64) uint64 {
	return b / 1024 / 1024
}
