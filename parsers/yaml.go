package parsers

import (
	"errors"

	"github.com/minskylab/tisp"
	"gopkg.in/yaml.v2"
)

type Metadata map[string]interface{}

type yamlBase struct {
	APIVersion string `yaml:"apiVersion"`
	Kind Kind `yaml:"kind"`
	Metadata Metadata `yaml:"metadata"`
	Spec interface{} `yaml:"spec"`
}

type Kind string

const Project Kind = "Project"
const Stage Kind = "Stage"
const Task Kind = "Task"
const Resource Kind = "Resource"

const UndefinedKind Kind = "Undefined"


type Bar struct {}

func readYAML(data []byte) (*yamlBase, error) {
	base := new(yamlBase)
	if err := yaml.Unmarshal(data, base); err != nil {
		return nil, err
	}
	return base, nil
}

func readProjectFromYAML(data []byte) (*tisp.NewProjectInformation, error) {
	base, err := readYAML(data)
	if err != nil {
		return nil, err
	}

	if base.Kind != Project {
		return nil, errors.New("your yaml isn't a project draft")
	}

	return parseYAMLInterfaceToProject(base.Spec)
}

func readTaskFromYAML(data []byte) (*tisp.NewTaskInformation, error) {
	base, err := readYAML(data)
	if err != nil {
		return nil, err
	}

	if base.Kind != Task {
		return nil, errors.New("your yaml isn't a task draft")
	}

	return parseYAMLInterfaceToTask(base.Spec)
}

func readResourceFromYAML(data []byte) (*tisp.NewResourceInformation, error) {
	base, err := readYAML(data)
	if err != nil {
		return nil, err
	}

	if base.Kind != Resource {
		return nil, errors.New("your yaml isn't a task draft")
	}

	return parseYAMLInterfaceToResource(base.Spec)
}

// TODO (bregydoc): to complete the stage level...