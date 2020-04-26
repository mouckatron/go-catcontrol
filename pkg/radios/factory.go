package radios

import (
	"errors"
)

func RadioFactory(radioName string) (*Radio, error) {

	switch radioName {
	case "yaesuft891":
		return NewYaesuFT891(), nil

	default:
		return nil, errors.New("Invalid Radio type")
	}
}

type RadioType struct {
	Name string
	Type string
}

var RadioTypes = []RadioType {
	RadioType{"Yaesu FT-891", "yaesuft891"},
	}
