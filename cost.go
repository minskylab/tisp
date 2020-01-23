package tisp

import (
	"errors"
	"strings"
)

type Cost struct {
	Units string
	Value float64
}

func (c *Cost) IsRatio() bool {
	return strings.Contains(c.Units, "/")
}

func (c *Cost) Upper() (string, error) {
	if !c.IsRatio() {
		return "", errors.New("the unit isn't a ratio, it's fixed")
	}

	chunks := strings.Split(c.Units, "/")
	return chunks[0], nil
}

func (c *Cost) Lower() (string, error) {
	if !c.IsRatio() {
		return "", errors.New("the unit isn't a ratio, it's fixed")
	}

	chunks := strings.Split(c.Units, "/")
	return chunks[1], nil
}

func (c *Cost) By() (string, error) {
	if !c.IsRatio() {
		return "", errors.New("the unit isn't a ratio, it's fixed")
	}

	chunks := strings.Split(c.Units, "/")
	return chunks[1], nil
}

func (c *Cost) IsValid() bool {
	// TODO: Add standard units validation ($, h, PEN, USD, etc...)
	return c.Units != ""
}
