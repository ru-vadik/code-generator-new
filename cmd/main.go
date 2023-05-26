package main

import (
	"cg-new/internal"
)

func main() {
	config := internal.NewConfig()
	randomCode := internal.NewRandomCode(config)
	planFile := internal.NewPlanFile(config)

	planFile.C <- randomCode.GetCode()
	planFile.C <- randomCode.GetCode()
	planFile.C <- randomCode.GetCode()
	planFile.C <- randomCode.GetCode()
	planFile.C <- randomCode.GetCode()

	close(planFile.C)
	config.Save()
}
