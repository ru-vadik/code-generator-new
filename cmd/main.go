package main

import (
	"cg-new/internal"
)

func main() {
	config := internal.NewConfig()
	planFile := internal.NewPlanFile(config)
	binaryTree := internal.NewBinaryTree(config, planFile)
	binaryTree.Start()

	config.Save()
}
