package motherboard

import (
	"os"
	"strings"
	"time"
)

type Bios struct {
	Brand   string
	Version string
	Date    time.Time
}

func GetBiosInfo() *Bios {
	brand, err := os.ReadFile("/sys/class/dmi/id/bios_vendor")
	if err != nil {
		panic(err)
	}

	version, err := os.ReadFile("/sys/class/dmi/id/bios_version")
	if err != nil {
		panic(err)
	}

	dateBytes, err := os.ReadFile("/sys/class/dmi/id/bios_date")
	if err != nil {
		panic(err)
	}

	dateStr := strings.TrimSpace(string(dateBytes))

	date, err := time.Parse("02/01/2006", dateStr)
	if err != nil {
		panic(err)
	}

	return &Bios{
		Brand:   string(brand),
		Version: string(version),
		Date:    date,
	}

}
