package main

import (
	"fmt"

	"cg-new/internal"
)

func main() {
	config := internal.NewConfig()

	fmt.Println(config)
	config.Save()
}
