package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"
)

func (db *StateManager) constructTaskTree(task tisp.NewTaskInformation) (*p.Task, error) {
	var leader *p.ResourceCreateOneWithoutTaskLeaderOfInput
	if task.Leader != nil {
		leader = &p.ResourceCreateOneWithoutTaskLeaderOfInput{
			Connect: &p.ResourceWhereUniqueInput{
				ID: p.Str(*task.Leader),
			},
		}
	}

	var resourcesInput *p.ResourceCreateManyWithoutWorkingOnInput
	if task.Resources != nil {
		resources := make([]p.ResourceWhereUniqueInput, 0)
		for _, id := range *task.Resources {
			resources = append(resources, p.ResourceWhereUniqueInput{
				ID: p.Str(id),
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
