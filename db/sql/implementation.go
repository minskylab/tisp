package sql

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/minskylab/tisp"
)

type DB struct {
	db *gorm.DB
}

func NewDB(path string) (*DB, error) {
	d, err := initDB(path)
	if err != nil {
		return nil, err
	}

	return &DB{db: d}, nil
}

func (db *DB) RegisterNewProject(info tisp.NewProjectInformation) (*tisp.Project, error) {
	p := tisp.Project{
		ID:        tisp.GenerateNewID(),
		Selector:  info.Selector,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      info.Name,
		ClientID:  info.ClientID,
		Tasks:     []tisp.Task{},
	}

	if db.db.Create(&p); &p == nil {
		return nil, errors.New("cannot perform the creation of this project")
	}

	return &p, nil
}

func (db *DB) RegisterNewResource(newRes tisp.NewResourceInformation) (*tisp.Resource, error) {
	r := tisp.Resource{
		ID:                       tisp.GenerateNewID(),
		Selector:                 newRes.Selector,
		CreatedAt:                time.Now(),
		UpdatedAt:                time.Now(),
		Name:                     newRes.Name,
		MainType:                 newRes.MainType,
		Cost:                     newRes.Cost,
		Types:                    []tisp.ResourceType{},
		WorkingOnAsProjectLeader: []tisp.Project{},
		WorkingOnAsTaskLeader:    []tisp.Task{},
		WorkingOn:                []tisp.Task{},
	}

	if newRes.Types != nil {
		for _, t := range *newRes.Types {
			r.Types = append(r.Types, t)
		}
	}

	if newRes.Alias != nil {
		r.Alias = *newRes.Alias
	}

	if newRes.Experience != nil {
		r.Experience = *newRes.Experience
	}

	if db.db.Create(&r); &r == nil {
		return nil, errors.New("cannot perform the creation of this resource")
	}

	return &r, nil
}

func (db *DB) AddSubTaskToTask(taskID tisp.ID, task tisp.NewTaskInformation) (*tisp.Task, error) {
	parentTask, err := db.GetTask(taskID)
	if err != nil {
		return nil, err
	}

	t := tisp.Task{
		ID:            tisp.GenerateNewID(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Title:         task.Title,
		Selector:      task.Selector,
		Resources:     []tisp.Resource{},
		State:         tisp.TaskState(""), // TODO(bregydoc) review
		Children:      []tisp.Task{},
		ProjectParent: "",
		ParentTask:    taskID,
	}

	if task.Leader != nil {
		t.LeaderID = *task.Leader
	}

	if task.Description != nil {
		t.Description = *task.Description
	}

	if task.Children != nil {
		for _, child := range *task.Children {
			if _, err = db.AddSubTaskToTask(parentTask.ID, child); err != nil {
				return nil, err
			}
		}
	}

	if db.db.Create(&t); &t == nil {
		return nil, errors.New("cannot perform the creation of this task")
	}

	newChildren := append(parentTask.Children, t)

	db.db.Model(parentTask).Update("Children", newChildren)

	return parentTask, nil
}

func (db *DB) AddTaskToProject(projectID tisp.ID, task tisp.NewTaskInformation) (*tisp.Project, error) {
	project, err := db.GetProject(projectID)
	if err != nil {
		return nil, err
	}

	t := &tisp.Task{
		ID:            tisp.GenerateNewID(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
		Title:         task.Title,
		Selector:      task.Selector,
		Resources:     []tisp.Resource{},
		State:         tisp.TaskState(""), // TODO(bregydoc) review
		Children:      []tisp.Task{},
		ProjectParent: project.ID,
	}

	if task.Leader != nil {
		t.LeaderID = *task.Leader
	}

	if task.Description != nil {
		t.Description = *task.Description
	}

	if task.Resources != nil {
		for _, r := range *task.Resources {
			res, err := db.GetResource(r)
			if err != nil {
				return nil, err
			}
			t.Resources = append(t.Resources, *res)
		}
	}

	if db.db.Create(&t); &t == nil {
		return nil, errors.New("cannot perform the creation of this task")
	}

	rootTaskID := t.ID

	if task.Children != nil {
		for _, child := range *task.Children {
			if t, err = db.AddSubTaskToTask(rootTaskID, child); err != nil {
				return nil, err
			}
		}
	}

	db.db.Model(project).Update("Tasks", append(project.Tasks, *t))

	return project, nil
}

func (db *DB) AddResourceToTask(taskID tisp.ID, res tisp.NewResourceInformation) (*tisp.Task, error) {
	panic("implement me")
}

func (db *DB) GetProject(projectID tisp.ID) (*tisp.Project, error) {
	project := new(tisp.Project)
	if db.db.First(project, projectID); project.ID == "" {
		return nil, errors.New("project with id '" + string(projectID) + "' not found")
	}

	return project, nil
}

func (db *DB) GetProjectBySelector(selector string) (*tisp.Project, error) {
	project := new(tisp.Project)
	if db.db.Where("selector = ?", selector); project.ID == "" {
		return nil, errors.New("project with selector '" + selector + "' not found")
	}

	return project, nil
}

func (db *DB) GetResource(resID tisp.ID) (*tisp.Resource, error) {
	resource := new(tisp.Resource)
	if db.db.First(resource, resID); resource.ID == "" {
		return nil, errors.New("resource with id '" + string(resID) + "' not found")
	}

	return resource, nil
}

func (db *DB) GetResourceBySelector(selector string) (*tisp.Resource, error) {
	resource := new(tisp.Resource)
	if db.db.Where("selector = ?", selector); resource.ID == "" {
		return nil, errors.New("resource with selector '" + selector + "' not found")
	}

	return resource, nil
}

func (db *DB) GetTask(taskID tisp.ID) (*tisp.Task, error) {
	task := new(tisp.Task)
	if db.db.First(task, taskID); task.ID == "" {
		return nil, errors.New("task with id '" + string(taskID) + "' not found")
	}

	return task, nil
}

func (db *DB) GetTaskBySelector(selector string) (*tisp.Task, error) {
	task := new(tisp.Task)
	if db.db.Where("selector = ?", selector); task.ID == "" {
		return nil, errors.New("task with selector '" + selector + "' not found")
	}

	return task, nil
}
