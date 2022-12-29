package presets

import (
	"fmt"
)

type PresetAlreadyExists struct {
	preset string
	base   bool
}

func (p *PresetAlreadyExists) Error() string {
	mssg := fmt.Sprintf("env preset '%s' already exists", p.preset)
	if p.base {
		mssg = "base " + mssg
	}
	return mssg
}

type PresetDoesNotExist struct {
	preset string
	base   bool
}

func (p *PresetDoesNotExist) Error() string {
	mssg := fmt.Sprintf("env preset '%s' does not exist", p.preset)
	if p.base {
		mssg = "base " + mssg
	}
	return mssg
}
