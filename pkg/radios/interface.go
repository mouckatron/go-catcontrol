package radios

type RadioBasic interface {
	sendCommand(command string, terminator string) []byte
	PowerState()
	PowerOn()
	PowerOff()
	Settings()
}
