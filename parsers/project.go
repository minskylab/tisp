package parsers

import (
	"errors"

	"github.com/minskylab/tisp"
)

func parseYAMLInterfaceToProject(data interface{}) (*tisp.NewProjectInformation, error) {
	// log.Info(reflect.TypeOf(data))
	switch data.(type) {
	case map[interface{}]interface{}:
		project := data.(map[string]interface{})

		switch {
		case len(project) > 2 && len(project) <= 3:
			name, _ := project["name"].(string)
			selector, _ := project["selector"].(string)
			partnerID, _ := project["partner"].(string)

			var sel *string
			if selector != "" {
				sel = &selector
			}
			return &tisp.NewProjectInformation{
				Name:      name,
				Selector:  sel,
				PartnerID: partnerID,
			}, nil
		}

		// TODO: Add more cases to enhance the resilience of the ctl
	}

	return nil, errors.New("invalid type, you need a map")
}
