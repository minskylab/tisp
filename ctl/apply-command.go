package main

import (
	"github.com/k0kubun/pp"
	"github.com/minskylab/tisp"
	"github.com/minskylab/tisp/parsers"
)

func (ctl *Control) apply(yamlFile string) error {
	kind, metadata, value, err := parsers.ReadFile(yamlFile)
	if err != nil {
		return err
	}

	switch kind {
	case parsers.Project:
		p, err := ctl.Machine.State.RegisterNewProject(value.(tisp.NewProjectInformation))
		if err != nil {
			return err
		}
		pp.Println(p)
	case parsers.Stage:
		// TODO...
	case parsers.Task:
		projectID, _ := metadata["project"].(string)
		task, err := ctl.Machine.State.AddTaskToProject(projectID, value.(tisp.NewTaskInformation))
		if err != nil {
			return err
		}
		pp.Println(task)
	case parsers.Resource:
		res, err := ctl.Machine.State.RegisterNewResource(value.(tisp.NewResourceInformation))
		if err != nil {
			return err
		}
		pp.Println(res)
	}

	return nil
}