package prisma

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/dbclient"
)

type DB struct {
	client *p.Client
}

func NewDB() (*DB, error) {
	return &DB{
		client: p.New(&p.Options{
			Endpoint: "http://localhost:4466",
			Secret:   "mysecret42",
		}),
	}, nil
}

func (db *DB) RegisterNewProject(info tisp.NewProjectInformation) (*p.Project, error) {
	return db.client.CreateProject(p.ProjectCreateInput{
		Selector: info.Selector,
		Name:     info.Name,
		Client: p.CompanyCreateOneWithoutProjectsInput{
			Connect: &p.CompanyWhereUniqueInput{
				ID: p.Str(string(info.ClientID)),
			},
		},
	}).Exec(context.Background())
}

func (db *DB) RegisterNewResource(newRes tisp.NewResourceInformation) (*p.Resource, error) {
	types := new(p.ResourceCreatetypesInput)
	if newRes.Types != nil {
		for _, t := range *newRes.Types {
			types.Set = append(types.Set, t)
		}
	}

	return db.client.CreateResource(p.ResourceCreateInput{
		Selector:   newRes.Selector,
		Name:       newRes.Name,
		MainType:   newRes.MainType,
		Types:      types,
		Alias:      newRes.Alias,
		Experience: newRes.Experience,
		Cost: p.CostCreateOneInput{
			Create: &p.CostCreateInput{
				Units: newRes.Cost.Units,
				Value: newRes.Cost.Value,
			},
		},
	}).Exec(context.Background())
}

func (db *DB) AddTaskToProject(projectID tisp.ID, task tisp.NewTaskInformation) (*p.Project, error) {
	rootTask, err := db.constructTaskTree(task)
	if err != nil {
		return nil, err
	}

	rootTask, err = db.client.UpdateTask(p.TaskUpdateParams{
		Data: p.TaskUpdateInput{
			Project: &p.ProjectUpdateOneWithoutTasksInput{
				Connect: &p.ProjectWhereUniqueInput{
					ID: p.Str(string(projectID)),
				},
			},
		},
		Where: p.TaskWhereUniqueInput{ID: p.Str(rootTask.ID)},
	}).Exec(context.Background())

	if err != nil {
		return nil, err
	}

	return db.client.UpdateProject(p.ProjectUpdateParams{
		Where: p.ProjectWhereUniqueInput{
			ID: p.Str(string(projectID)),
		},
		Data: p.ProjectUpdateInput{
			Tasks: &p.TaskUpdateManyWithoutProjectInput{
				Connect: []p.TaskWhereUniqueInput{
					{ID: p.Str(rootTask.ID)},
				},
			},
		},
	}).Exec(context.Background())
}

func (db *DB) constructTaskTree(task tisp.NewTaskInformation) (*p.Task, error) {
	var leader *p.ResourceCreateOneWithoutTaskLeaderOfInput
	if task.Leader != nil {
		leader = &p.ResourceCreateOneWithoutTaskLeaderOfInput{
			Connect: &p.ResourceWhereUniqueInput{
				ID: p.Str(string(*task.Leader)),
			},
		}
	}

	var resourcesInput *p.ResourceCreateManyWithoutWorkingOnInput
	if task.Resources != nil {
		resources := make([]p.ResourceWhereUniqueInput, 0)
		for _, id := range *task.Resources {
			resources = append(resources, p.ResourceWhereUniqueInput{
				ID: p.Str(string(id)),
			})
		}
		resourcesInput = &p.ResourceCreateManyWithoutWorkingOnInput{
			Connect: resources,
		}
	}

	newTask, err := db.client.CreateTask(p.TaskCreateInput{
		Selector:    task.Selector,
		Title:       task.Title,
		Description: task.Description,
		Leader:      leader,
		Resources:   resourcesInput,
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	if task.Children != nil {
		for _, c := range *task.Children {
			var children *p.TaskUpdateManyWithoutParentInput
			if c.Children != nil {
				newChild, err := db.constructTaskTree(c)
				if err != nil {
					return nil, err
				}
				children = &p.TaskUpdateManyWithoutParentInput{
					Connect: []p.TaskWhereUniqueInput{
						{ID: p.Str(newChild.ID)},
					},
				}
			}

			_, err := db.client.UpdateTask(p.TaskUpdateParams{
				Where: p.TaskWhereUniqueInput{ID: p.Str(newTask.ID)},
				Data: p.TaskUpdateInput{
					Children: children,
				},
			}).Exec(context.Background())
			if err != nil {
				return nil, err
			}
		}
	}

	return newTask, nil
}

func (db *DB) AddSubTaskToTask(taskID tisp.ID, task tisp.NewTaskInformation) (*p.Task, error) {
	newTask, err := db.constructTaskTree(task)
	if err != nil {
		return nil, err
	}

	return db.client.UpdateTask(p.TaskUpdateParams{
		Data: p.TaskUpdateInput{
			Children: &p.TaskUpdateManyWithoutParentInput{
				Connect: []p.TaskWhereUniqueInput{
					{ID: p.Str(newTask.ID)},
				},
			},
		},
		Where: p.TaskWhereUniqueInput{
			ID: p.Str(string(taskID)),
		},
	}).Exec(context.Background())
}

func (db *DB) AddResourceToTask(taskID tisp.ID, res tisp.NewResourceInformation) (*p.Task, error) {
	types := new(p.ResourceCreatetypesInput)
	if res.Types != nil {
		for _, t := range *res.Types {
			types.Set = append(types.Set, t)
		}
	}

	return db.client.UpdateTask(p.TaskUpdateParams{
		Where: p.TaskWhereUniqueInput{
			ID: p.Str(string(taskID)),
		},
		Data: p.TaskUpdateInput{
			Resources: &p.ResourceUpdateManyWithoutWorkingOnInput{
				Create: []p.ResourceCreateWithoutWorkingOnInput{
					{
						Selector:   res.Selector,
						Name:       res.Name,
						MainType:   res.MainType,
						Types:      types,
						Alias:      res.Alias,
						Experience: res.Experience,
						Cost: p.CostCreateOneInput{
							Create: &p.CostCreateInput{
								Units: res.Cost.Units,
								Value: res.Cost.Value,
							},
						},
					},
				},
			},
		},
	}).Exec(context.Background())
}

func (db *DB) GetProject(projectID tisp.ID) (*p.Project, error) {
	return db.client.Project(p.ProjectWhereUniqueInput{
		ID: p.Str(string(projectID)),
	}).Exec(context.Background())
}

func (db *DB) GetProjectBySelector(selector string) (*p.Project, error) {
	return db.client.Project(p.ProjectWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *DB) GetResource(resID tisp.ID) (*p.Resource, error) {
	return db.client.Resource(p.ResourceWhereUniqueInput{
		ID: p.Str(string(resID)),
	}).Exec(context.Background())
}

func (db *DB) GetResourceBySelector(selector string) (*p.Resource, error) {
	return db.client.Resource(p.ResourceWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *DB) GetTask(taskID tisp.ID) (*p.Task, error) {
	return db.client.Task(p.TaskWhereUniqueInput{
		ID: p.Str(string(taskID)),
	}).Exec(context.Background())
}

func (db *DB) GetTaskBySelector(selector string) (*p.Task, error) {
	return db.client.Task(p.TaskWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}
