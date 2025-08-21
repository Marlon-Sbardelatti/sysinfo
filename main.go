package main

import (
	"fmt"
	"systop/cpu"
	// "systop/motherboard"
)

func main()  {
	// mobo := motherboard.GetMoboInfo()

	// bios := motherboard.GetBiosInfo()	
	// fmt.Println()
	// fmt.Printf("Motherboard: %v", mobo)
	// fmt.Println()
	// fmt.Println()
	// fmt.Printf("Bios: %v", bios)
	// fmt.Println()
	// fmt.Println()
	// fmt.Println()


	cpuStatic := cpu.GetCpuStaticInfo()
	fmt.Printf("%+v", cpuStatic)
}
