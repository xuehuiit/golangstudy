package study

import (
	"fmt"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	//"github.com/shirou/gopsutil/process"
)

func Getos() {

	fmt.Println(runtime.GOARCH)
	fmt.Println(runtime.GOOS)

}

func GetCurrMem() {

	v, _ := mem.VirtualMemory()

	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

}

func GetCurrPro() {

	//v, _ := process.

	fmt.Printf("Total: %v, Free:%v, UsedPercent:%f%%\n", v.Total, v.Free, v.UsedPercent)

	// convert to JSON. String() is also implemented
	fmt.Println(v)

}
