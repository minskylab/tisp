package parsers

import (
	"errors"

	"github.com/minskylab/tisp"
)

func parseYAMLInterfaceToTask(data interface{}) (*tisp.NewTaskInformation, error) {
	switch data.(type) {
	case map[string]interface{}:
		project := data.(map[string]interface{})

		title, _ := project["title"].(string)
		description, _ := project["description"].(string)

		var sel *string
		selector, _ := project["selector"].(string)
		if selector != "" {
			sel = &selector
		}

		var leader *string
		leaderID, _ := project["leader"].(string)
		if leaderID != "" {
			leader = &leaderID
		}

		var subTasks *[]tisp.NewTaskInformation
		tasks, withSubTasks := project["tasks"].([]interface{})
		if withSubTasks {
			for _, t := range tasks {
				newTask, err := parseYAMLInterfaceToTask(t)
				if err != nil {
					return nil, err
				}
				if subTasks == nil {
					subTasks = &[]tisp.NewTaskInformation{}
				}
				*subTasks = append(*subTasks, *newTask)
			}
		}

		var resources *[]string
		res, withResources := project["resources"].([]string)
		if withResources {
			resources = &res
		}

		return &tisp.NewTaskInformation{
			Title:       title,
			Description: description,
			Selector:    sel,
			Resources:   resources,
			Leader:      leader,
			Children:    subTasks,
		}, nil

		// TODO: Add more cases to enhance the resilience of the ctl
	}

	return nil, errors.New("invalid type, you need a map")
}