package parsers

import (
	"errors"
	"io/ioutil"

	"github.com/minskylab/tisp"
)

func ReadFile(filepath string) (Kind, Metadata, interface{}, error) {
	data, err :=ioutil.ReadFile(filepath)
	if err != nil {
		return UndefinedKind, nil, nil, err
	}

	base, err := readYAML(data)
	if err != nil {
		return UndefinedKind, nil, nil, err
	}

	if err := tisp.IsCorrectVersion(base.APIVersion); err != nil {
		return UndefinedKind, nil, nil, err
	}

	switch base.Kind {
	case Project:
		project, err := readProjectFromYAML(data)
		return base.Kind, base.Metadata, project, err
	case Stage:
		// stage, err := readSt
		// TODO...
	case Task:
		task, err := readTaskFromYAML(data)
		return base.Kind, base.Metadata, task, err
	case Resource:
		resource, err := readResourceFromYAML(data)
		return base.Kind, base.Metadata, resource, err
	}

	return UndefinedKind, nil, nil, errors.New("invalid kind, please use one valid")
}