package main

import (
	"cg-new/internal/cg"
)

func main() {
	config := cg.NewConfig()
	planFile := cg.NewPlanFile(config)
	binaryTree := cg.NewBinaryTree(config, planFile)
	binaryTree.Start()

	config.Save()
}
