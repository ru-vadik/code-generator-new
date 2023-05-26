package main

import (
	"cg-new/internal"
	"fmt"
)

func main() {
	config := internal.NewConfig()
	planFile := internal.NewPlanFile(config)
	planFile.C <- "a"
	planFile.C <- "b"
	close(planFile.C)
	config.Save()

	fmt.Println("OK")
}
