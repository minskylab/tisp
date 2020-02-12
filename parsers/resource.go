package parsers

import (
	"errors"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"
)

func parseYAMLInterfaceToResource(data interface{}) (*tisp.NewResourceInformation, error) {
	switch data.(type) {
	case map[string]interface{}:
		task := data.(map[string]interface{})

		name, _ := task["name"].(string)
		cost, _ := task["cost"].(tisp.Cost)
		mainType, _ := task["mainType"].(string)
		experience, _ := task["experience"].(float64)

		var sel *string
		selector, _ := task["selector"].(string)
		if selector != "" {
			sel = &selector
		}

		var alias *string
		al, _ := task["alias"].(string)
		if al != "" {
			alias = &al
		}

		var types *[]p.ResourceType
		t, ok := task["types"].([]p.ResourceType)
		if ok {
			types = &t
		}

		return &tisp.NewResourceInformation{
			Name:       name,
			Cost:       cost,
			MainType:   p.ResourceType(mainType),
			Experience: experience,
			Selector:   sel,
			Alias:      alias,
			Types:      types,
		}, nil
	}

	return nil, errors.New("invalid data type, please write a correct resource description")
}
