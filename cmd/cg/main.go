package main

import (
	"cg-new/internal/cg"
	"flag"
	"fmt"
	"os"
	"runtime/pprof"
	"strconv"
	"time"
)

func main() {
	startTime := time.Now()

	cpuProfileFlag := flag.Bool("cpuProfile", false, "")
	flag.Parse()
	if *cpuProfileFlag {
		cpuProfileFile, err := os.Create("cpuProfile.prof")
		if err != nil {
			panic(err)
		}
		err = pprof.StartCPUProfile(cpuProfileFile)
		if err != nil {
			panic(err)
		}
		defer pprof.StopCPUProfile()
	}

	config := cg.NewConfig()
	planFile := cg.NewPlanFile(config)
	binaryTree := cg.NewBinaryTree(config, planFile)
	binaryTree.Start()

	config.Save()
	fmt.Println("File: " + planFile.FileName)
	fmt.Println("Qty: " + strconv.FormatInt(int64(config.Qty), 10))
	fmt.Println("Time: " + time.Since(startTime).String())
}
