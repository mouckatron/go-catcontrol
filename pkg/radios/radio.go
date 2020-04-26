package radios

import (
	"errors"
	"log"
	"time"

	"github.com/tarm/serial"
)

type Radio struct {
	Name string
	Manufacturer string
	Model string
	Commands []RadioCommand //TODO This should be a map so we can wrap this base list in an interfacey thing
	commandTerminator string
}

func (r Radio) sendCommand(command string) (result []byte) {
	timeout, _ := time.ParseDuration("1s")
	c := &serial.Config{
		Name: "/dev/ttyUSB0",
		Baud: 38400,
		ReadTimeout: timeout,
	}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Writing %s", command)
	_, err = s.Write([]byte(command))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Reading response")
	buf := make([]byte, 1)
	for {
		_, err := s.Read(buf)
		if err != nil {
			//log.Fatal(err)
			return nil
		}

		result = append(result, buf...)
		if string(buf) == r.commandTerminator {
			break
		}
	}
	return
}

func (r *Radio) Settings() {
	for _, c := range r.Commands {
		log.Printf("%s", c.toString())
	}
}

func (r *Radio) GetSetting(setting string) (string, error) {
	for _, cmd := range r.Commands {
		if cmd.Command() == setting {
			return cmd.get()
		}
	}
	return "", errors.New("Command not found")
}

func (r *Radio) SetSetting(setting string, value string) (error) {
	for _, cmd := range r.Commands {
		if cmd.Command() == setting {
			return cmd.set(value)
		}
	}
	return errors.New("Command not found")
}
