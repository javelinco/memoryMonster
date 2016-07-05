package main

import (
        "log"
        "runtime"
	"time"
	"os"
	"strconv"
)

func bigBytes(allocSize int) *[]byte {
        s := make([]byte, allocSize)
        return &s
}

func memReport() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println("")
	log.Println("mem.Alloc = " + strconv.FormatUint(mem.Alloc, 10))
	log.Println("mem.TotalAlloc = " + strconv.FormatUint(mem.TotalAlloc, 10))
	log.Println("mem.HeapAlloc = " + strconv.FormatUint(mem.HeapAlloc, 10))
	log.Println("mem.HeapSys = " + strconv.FormatUint(mem.HeapSys, 10))
}
 
func main() {
	memReport()

	var allocSize int 
	allocSize, error := strconv.Atoi(os.Args[1])
	if error != nil {
		allocSize = 1000000
	}
        for {
                s := bigBytes(allocSize)
                if s == nil {
                        log.Println("oh noes")
                } else {
			memReport()
		}
		time.Sleep(1 * time.Second)
        }
}
