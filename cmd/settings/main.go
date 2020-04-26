package main

import (
	"flag"
	"log"

	"github.com/mouckatron/go-catcontrol/pkg/radios"
)

func main() {

	radio := flag.String("radio", "yaesuft891", "Radio type")
	var setting string = ""
	var settingValue string = ""
	flag.Parse()

	if flag.NArg() > 0 {
		setting = flag.Arg(0)
	}

	if flag.NArg() > 1 {
		settingValue = flag.Arg(1)
	}

	log.Printf("Radio: %s", *radio)

	r, err := radios.RadioFactory(*radio)
	if err != nil {
		log.Fatal(err)
	}

	if setting != "" {
		if settingValue != "" {
			err := r.SetSetting(setting, settingValue)
			if err != nil {
				log.Fatal(err)
			}

		} else {
			response, err := r.GetSetting(setting)
			if err != nil {
				log.Fatal(err)
			}
			print(response)
		}

	} else {
		r.Settings()
	}
}
