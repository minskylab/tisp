package parsers

import (
	"errors"

	"gopkg.in/yaml.v2"
)

type yamlBase struct {
	APIVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata interface{} `yaml:"metadata"`
	Spec interface{} `yaml:"spec"`
}

type Kind string

const Project Kind = "project"
const Stage Kind = "stage"
const Task Kind = "task"
const Resource Kind = "resource"

const UndefinedKind Kind = "undefined"

func indentifyAndVerifyYAML(data []byte) (Kind, error) {
	base := new(yamlBase)
	if err := yaml.Unmarshal(data, base); err != nil {
		return UndefinedKind, err
	}

	switch base.Kind {
	case Project:

	case Stage:
	case Task:
	case Resource:
	default:
		return UndefinedKind, errors.New("invalid kind")
	}

}