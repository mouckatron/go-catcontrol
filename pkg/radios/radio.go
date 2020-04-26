package radios

import (
	"bytes"
	"errors"
	"fmt"
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

type RadioCommand interface {
	toString() string
	get() (string, error)
	set(value string)
}

// Information about the command
type RadioCommandOptions struct {
	radio *Radio
	name string
	command string
	read bool
	write bool
	readCommand string
	writeCommand string
	options []RadioCommandOption
}

type RadioCommandOption struct {
	name string
	value string
}

func (rc RadioCommandOptions) toString() string {

	var output bytes.Buffer
	output.WriteString(rc.name)
	output.WriteString(":")
	result, _ := rc.get()
	output.WriteString(result)
	output.WriteString(" options[")

	for _, x := range rc.options {
		output.WriteString(x.name)
		output.WriteString(":")
		output.WriteString(x.value)
		output.WriteString(",")
	}

	return fmt.Sprintf("%s]", string(bytes.TrimSuffix(output.Bytes(), []byte{','})))
}

func (rc RadioCommandOptions) set(value string) {

}

func (rc RadioCommandOptions) get() (string, error) {
	if rc.read {
		response := string(rc.radio.sendCommand(rc.readCommand))
		return response, nil
	}
	return "", errors.New("Command does not support read")
}

type RadioCommandString struct {
	radio *Radio
	name string
	command string
	read bool
	write bool
	readCommand string
	writeCommand string
}

func (rc RadioCommandString) toString() string {

	result, _ := rc.get()
	return fmt.Sprintf("%s: %s", rc.name, result)
}

func (rc RadioCommandString) set(value string) {

}

func (rc RadioCommandString) get() (string, error) {
	if rc.read {
		response := string(rc.radio.sendCommand(rc.readCommand))
		return response, nil
	}
	return "", errors.New("Command does not support read")
}


// // run the command against the radio
// type RadioSetting struct {
// 	RadioCommand
// }

// func (r RadioSetting) getSetting(*r RadioBasic)
