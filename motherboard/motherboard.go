package motherboard

import "os"

type Motherboard struct {
	Manufacturer string
	Model        string
	Revision     string
}

func GetMoboInfo() *Motherboard {
	manufacturer, err := os.ReadFile("/sys/class/dmi/id/board_vendor")
	if err != nil {
		panic(err)
	}

	model, err := os.ReadFile("/sys/class/dmi/id/board_name")
	if err != nil {
		panic(err)
	}

	revision, err := os.ReadFile("/sys/class/dmi/id/board_version")
	if err != nil {
		panic(err)
	}

	return &Motherboard{
		Manufacturer: string(manufacturer),
		Model: string(model),
		Revision: string(revision),
	}
}
