package main

import (
	"flag"
	"log"
	"os"

	"github.com/mouckatron/go-catcontrol/pkg/radios"
)

func main() {

	radio := flag.String("radio", "yaesuft891", "Radio type")
	action := os.Args[len(os.Args)-1]

	flag.Parse()

	log.Printf("Radio: %s", *radio)

	r, err := radios.RadioFactory(*radio)

	if err != nil {
		log.Fatal(err)
	}

	log.Print(action)

	switch action {
	case "on":
		r.PowerOn()
		r.PowerOn()
	case "off":
		r.PowerOff()
	default:
		r.PowerState()
	}
}
