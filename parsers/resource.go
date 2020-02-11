package parsers

import (
	"errors"

	"github.com/minskylab/tisp"
)

func parseYAMLInterfaceToResource(data interface{}) (*tisp.NewResourceInformation, error) {

	switch data.(type) {
	case map[string]interface{}:
		task := data.(map[string]interface{})

		name, _ := task["name"].(string)

		return &tisp.NewResourceInformation{
			Name:       name,
			Cost:       tisp.Cost{},
			MainType:   "",
			Experience: 0,
			Selector:   nil,
			Alias:      nil,
			Types:      nil,
		}, nil
	}

	return nil, errors.New("invalid data type, please write a correct resource description")
}
