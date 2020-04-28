package main

import (
	"flag"
	"log"

	"github.com/mouckatron/go-catcontrol/pkg/radios"
	"gopkg.in/ini.v1"
)

func main() {

	radio := flag.String("radio", "yaesuft891", "Radio type")
	file := flag.String("file", "", "Settings file to load")
	var setting string = ""
	var settingValue string = ""
	flag.Parse()

	log.Printf("Radio: %s", *radio)
	r := getRadio(*radio)

	if *file != "" {
		runCommandsFromFile(r, *file)
		return
	}

	if flag.NArg() > 0 {
		setting = flag.Arg(0)
	}

	if flag.NArg() > 1 {
		settingValue = flag.Arg(1)
	}

	if setting != "" {
		runSetting(r, setting, settingValue)
	} else {
		allSettings(r)
		r.Settings()
	}
}

func getRadio(radio string) (*radios.Radio) {
	r, err := radios.RadioFactory(radio)
	if err != nil {
		log.Fatal(err)
	}
	return r
}

func runCommandsFromFile(r *radios.Radio, filename string) {
	cfg, err := ini.Load(filename)
	if err != nil {
		log.Fatal("Could not read file", err)
	}

	for key, val := range cfg.Section("").KeysHash() {
		r.SetSetting(key, val)
	}

}

func runSetting(r *radios.Radio, setting string, settingValue string) {
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
}

func allSettings(r *radios.Radio) {
	r.Settings()
}
