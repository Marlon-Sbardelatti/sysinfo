package cpu

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type CpuStatic struct {
	VendorId     string
	ModelName    string
	Stepping     string
	Microcode    string
	Cores        int
	Threads      int
	Architecture string
}

func GetCpuStaticInfo() *CpuStatic {
	vendorId, modelName, stepping, microcode, cores := getProcCpuInfo("/proc/cpuinfo")
	threads := getTotalThreads("/sys/devices/system/cpu/online")
	// threads := getTotalThreads("/home/hetzwga/testecpu.txt")
	return &CpuStatic{
		VendorId:  vendorId,
		ModelName: modelName,
		Stepping:  stepping,
		Microcode: microcode,
		Cores:     cores,
		Threads:   threads,
	}
}

func getProcCpuInfo(path string) (string, string, string, string, int) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	var vendorId string
	var modelName string
	var stepping string
	var microcode string
	var cores int

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		} else if strings.HasPrefix(line, "vendor_id") {
			vendorId = strings.Split(line, ":")[1]
		} else if strings.HasPrefix(line, "model name") {
			modelName = strings.Split(line, ":")[1]
		} else if strings.HasPrefix(line, "stepping") {
			stepping = strings.Split(line, ":")[1]
		} else if strings.HasPrefix(line, "microcode") {
			microcode = strings.Split(line, ":")[1]
		} else if strings.HasPrefix(line, "cpu cores") {
			coresStr := strings.Split(line, ":")[1]
			cores, err = strconv.Atoi(strings.TrimSpace(coresStr))
			if err != nil {
				panic(err)
			}
		} else if strings.HasPrefix(line, "flags") {

		}

	}

	return vendorId, modelName, stepping, microcode, cores
}

func getTotalThreads(path string) int {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	threadStr := strings.TrimSpace(string(file))
	threadNumbers := strings.Split(threadStr, "-")

	var idx int
	if len(threadNumbers) == 1 {
		idx = 0
	} else {
		idx = 1
	}

	threads, err := strconv.Atoi(threadNumbers[idx])
	if err != nil {
		panic(err)
	}

	return threads + 1
}
