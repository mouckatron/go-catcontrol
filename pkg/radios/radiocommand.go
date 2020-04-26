package radios

import (
	"bytes"
	"errors"
	"fmt"
)

type RadioCommand interface {
	toString() string
	Name() (string)
	Command() (string)
	get() (string, error)
	set(value string) (error)
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

func (rc RadioCommandOptions) Command() string { return rc.command }
func (rc RadioCommandOptions) Name() string { return rc.name }
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

func (rc RadioCommandOptions) set(value string) (error) {
	if !rc.write {
		return errors.New("Command does not support write")
	}

	if !rc.isValidOption(value) {
		return errors.New("Invalid option to command")
	}

	rc.radio.sendCommand(fmt.Sprintf(rc.writeCommand, value))
	return nil
}

func (rc RadioCommandOptions) get() (string, error) {
	if !rc.read {
		return "", errors.New("Command does not support read")
	}

	response := string(rc.radio.sendCommand(rc.readCommand))
	return response, nil
}

func (rc RadioCommandOptions) isValidOption(value string) (bool) {

	for _, opt := range rc.options {
		if opt.value == value {
			return true
		}
	}

	return false
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

func (rc RadioCommandString) Command() string { return rc.command }
func (rc RadioCommandString) Name() string { return rc.name }
func (rc RadioCommandString) toString() string {

	result, _ := rc.get()
	return fmt.Sprintf("%s: %s", rc.name, result)
}

func (rc RadioCommandString) set(value string) (error) {
	if !rc.write {
		return errors.New("Command does not support write")
	}

	rc.radio.sendCommand(fmt.Sprintf(rc.writeCommand, value))
	return nil
}

func (rc RadioCommandString) get() (string, error) {
	if !rc.read {
		return "", errors.New("Command does not support read")
	}

	response := string(rc.radio.sendCommand(rc.readCommand))
	return response, nil
}
