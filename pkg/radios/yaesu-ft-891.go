package radios

func NewYaesuFT891() (*Radio) {
	r := &Radio{
		Name: "Yaesu FT-891",
		Manufacturer: "Yaesu",
		Model: "FT-891",
		commandTerminator: ";",
	}
	r.Commands = []RadioCommand{
		RadioCommandOptions{r, "Clarifier", "CF", true, true, "CF0;", "CF0%s0;",
			[]RadioCommandOption{RadioCommandOption{"On", "1"}, RadioCommandOption{"Off", "0"}}},
		RadioCommandOptions{r, "Power Switch", "PS", true, true, "PS;", "PS%s;",
			[]RadioCommandOption{RadioCommandOption{"On", "1"}, RadioCommandOption{"Off", "0"}}},
		RadioCommandString{r, "Frequency VFO-A", "FA", true, true, "FA;", "FA%s;"},
	}
    return r
}

//TODO These functions shouldn't exist here, they should be called externally
// func (r *YaesuFT891) PowerState() {
// 	response := r.sendCommand("PS;", r.terminator)

// 	log.Printf("%s", response)
// }

// func (r *YaesuFT891) PowerOn() {
// 	response := r.sendCommand("PS1;", r.terminator)
// 	log.Printf("%s", response)
// }

// func (r *YaesuFT891) PowerOff() {
// 	response := r.sendCommand("PS0;", r.terminator)
// 	log.Printf("%s", response)
// }
