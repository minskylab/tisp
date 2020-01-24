package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/dbclient"
)

func (db *StateManager) AddSubTaskToTask(taskID string, task tisp.NewTaskInformation) (*p.Task, error) {
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
			ID: p.Str(taskID),
		},
	}).Exec(context.Background())
}


func (db *StateManager) AddTaskToProject(projectID string, task tisp.NewTaskInformation) (*p.Project, error) {
	rootTask, err := db.constructTaskTree(task)
	if err != nil {
		return nil, err
	}

	rootTask, err = db.client.UpdateTask(p.TaskUpdateParams{
		Data: p.TaskUpdateInput{
			Project: &p.ProjectUpdateOneWithoutTasksInput{
				Connect: &p.ProjectWhereUniqueInput{
					ID: p.Str(projectID),
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
			ID: p.Str(projectID),
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