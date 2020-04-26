package main

import (
	"flag"
	"log"

	"github.com/mouckatron/go-catcontrol/pkg/radios"
)

func main() {

	radio := flag.String("radio", "yaesuft891", "Radio type")

	flag.Parse()

	log.Printf("Radio: %s", *radio)

	r, err := radios.RadioFactory(*radio)
	if err != nil {
		log.Fatal(err)
	}

	r.Settings()
}
