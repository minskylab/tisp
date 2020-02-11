package parsers

import "gopkg.in/yaml.v2"

type yamlBase struct {
	APIVersion string `yaml:"apiVersion"`
	Kind string `yaml:"kind"`
	Metadata interface{} `yaml:"metadata"`
	Spec interface{} `yaml:"spec"`
}

type Kind string

const Project Kind = "project"
const Resource Kind = "project"
const Resource Kind = "project"


func indentifyAndVerifyYAML(data []byte) (Kind, error) {
	base := new(yamlBase)
	if err := yaml.Unmarshal(data, base); err != nil {
		return err
	}


}