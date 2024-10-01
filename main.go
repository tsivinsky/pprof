package main

import (
	"flag"
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"strconv"
	"strings"
)

/*
real    0m15.695s
user    0m12.612s
sys     0m11.064s
*/
func foo() {
	s := ""
	for i := 0; i < 1_000_000; i++ {
		if i%5 == 0 {
			s += strconv.Itoa(i) + "\n"
		}
	}
	fmt.Println(s)
}

/*
real    0m0.833s
user    0m0.015s
sys     0m0.110s
*/
func bar() {
	r := []string{}
	for i := 0; i < 1_000_000; i++ {
		if i%5 == 0 {
			r = append(r, strconv.Itoa(i))
		}
	}
	fmt.Println(strings.Join(r, "\n"))
}

func main() {
	flag.Parse()

	funct := flag.Arg(0)
	if funct == "" {
		funct = "foo"
	}

	os.Remove("mem.pprof")

	f, err := os.OpenFile("mem.pprof", os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		panic(err)
	}
	defer pprof.StopCPUProfile()

	switch funct {
	case "foo":
		foo()
	case "bar":
		bar()
	default:
		panic("unknown function")
	}

	runtime.GC()
}
