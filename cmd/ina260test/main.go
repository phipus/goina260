package main

import (
	"flag"
	"log"
	"time"

	"github.com/d2r2/go-logger"
	ina260 "github.com/phipus/goina260"
)

var (
	flagAddress = flag.Int("address", 0x40, "Specify the i2c address (default = 0x40")
	flagBus     = flag.Int("bus", 0x01, "Specify the i2c bus")
)

func main() {
	log.SetPrefix("")
	logger.ChangePackageLogLevel("i2c", logger.WarnLevel)

	flag.Parse()
	sensor, err := ina260.New(uint8(*flagAddress), *flagBus)
	if err != nil {
		log.Fatal(err)
	}

	for {
		voltage, current, power, err := sensor.ReadData(true, true, true)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%f V    %f W    %f A", voltage, current, power)

		time.Sleep(500 * time.Millisecond)
	}
}
