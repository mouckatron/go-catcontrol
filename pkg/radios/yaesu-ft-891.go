package radios

func NewYaesuFT891() (*Radio) {
	r := &Radio{
		Name: "Yaesu FT-891",
		Manufacturer: "Yaesu",
		Model: "FT-891",
		commandTerminator: ";",
	}
	r.Commands = []RadioCommand{
// Option Commands
		RadioCommandOptions{r, "Band Select", "BS", false, true, "", "BS%s",
			[]RadioCommandOption{
				RadioCommandOption{"1.8MHz (160m)", "00"},
				RadioCommandOption{"3.5MHz (80m)", "01"},
				RadioCommandOption{"5MHz (60m)", "02"},
				RadioCommandOption{"7MHz (40m)", "03"},
				RadioCommandOption{"10MHz (30m)", "04"},
				RadioCommandOption{"14MHz (20m)", "05"},
				RadioCommandOption{"18MHz (17m)", "06"},
				RadioCommandOption{"21MHz (15m)", "07"},
				RadioCommandOption{"24.5MHz (12m)", "08"},
				RadioCommandOption{"28MHz (10m)", "09"},
				RadioCommandOption{"50MHz (6m)", "10"},
				RadioCommandOption{"GEN", "11"},
				RadioCommandOption{"MW", "12"},
			}},
		RadioCommandOptions{r, "Clarifier", "CF", true, true, "CF0;", "CF0%s0;",
			[]RadioCommandOption{RadioCommandOption{"On", "1"}, RadioCommandOption{"Off", "0"}}},
		RadioCommandOptions{r, "Operating Mode", "MD", true, true, "MD0;", "MD0%s;",
			[]RadioCommandOption{
				RadioCommandOption{"SSB (SSB BFO)", "1"},
				RadioCommandOption{"SSB (SSB BFO)", "2"},
				RadioCommandOption{"CW (CW BFO)", "3"},
				RadioCommandOption{"FM", "4"},
				RadioCommandOption{"AM", "5"},
				RadioCommandOption{"RTTY (RTTY BFO)", "6"},
				RadioCommandOption{"CW (CW BFO)", "7"},
				RadioCommandOption{"DATA (DATA BFO)", "8"},
				RadioCommandOption{"RTTY (RTTY BFO)", "9"},
				//RadioCommandOption{"-", "A"},
				RadioCommandOption{"FM-N", "B"},
				RadioCommandOption{"DATA (DATA BFO)", "C"},
				RadioCommandOption{"AM-N", "D"},
			}},
		RadioCommandOptions{r, "Power Switch", "PS", true, true, "PS;", "PS%s;",
			[]RadioCommandOption{RadioCommandOption{"On", "1"}, RadioCommandOption{"Off", "0"}}},
		RadioCommandOptions{r, "Scan", "SC", true, true, "SC;", "SC%s;",
			[]RadioCommandOption{
				RadioCommandOption{"Off", "0"},
				RadioCommandOption{"On (Down)", "1"},
				RadioCommandOption{"On (Up)", "2"},
			}},

// String Commands
		RadioCommandString{r, "Frequency VFO-A", "FA", true, true, "FA;", "FA%s;"},
		RadioCommandString{r, "Identification", "ID", true, false, "ID;", ""},
		RadioCommandString{r, "Information", "IF", true, false, "IF;", ""},
	}
    return r
}
