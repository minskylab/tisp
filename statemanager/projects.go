package statemanager

import (
	"context"
	"errors"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"
)

func (db *StateManager) RegisterNewProject(info tisp.NewProjectInformation) (*p.Project, error) {
	return db.client.CreateProject(p.ProjectCreateInput{
		Selector: info.Selector,
		Name:     info.Name,
		Partner: p.PartnerCreateOneWithoutProjectsInput{
			Connect: &p.PartnerWhereUniqueInput{
				ID: p.Str(info.PartnerID),
			},
		},
	}).Exec(context.Background())
}

func (db *StateManager) AddStageToProject(projectID string, newStage tisp.NewStage) (*p.Project, error) {
	stage, err := db.client.CreateStage(p.StageCreateInput{
		Name:     newStage.Name,
		Selector: newStage.Selector,
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	if newStage.Tasks != nil {
		for _, t := range *newStage.Tasks {
			if _, err := db.AddTaskToStage(projectID, stage.ID, t); err != nil {
				return nil, err
			}
		}
	}

	return db.client.UpdateProject(p.ProjectUpdateParams{
		Data: p.ProjectUpdateInput{
			Stages: &p.StageUpdateManyWithoutParentProjectInput{
				Connect: []p.StageWhereUniqueInput{
					{ID: p.Str(stage.ID)},
				},
			}},
		Where: p.ProjectWhereUniqueInput{ID: p.Str(projectID)},
	}).Exec(context.Background())

}

func (db *StateManager) AddTaskToStage(projectID, stageID string, task tisp.NewTaskInformation) (*p.Project, error) {
	stages, err := db.client.Project(
		p.ProjectWhereUniqueInput{
			ID: p.Str(projectID),
		}).Stages(&p.StagesParamsExec{
		Where: &p.StageWhereInput{ID: p.Str(stageID)},
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	if len(stages) == 0 {
		return nil, errors.New("your stageID not exists inside of your projectID")
	}

	t, err := db.constructTaskTree(task)
	if err != nil {
		return nil, err
	}

	_, err = db.client.UpdateStage(p.StageUpdateParams{
		Data: p.StageUpdateInput{
			Tasks: &p.TaskUpdateManyWithoutStageInput{
				Connect: []p.TaskWhereUniqueInput{
					{ID: p.Str(t.ID)},
				},
			},
		},
		Where: p.StageWhereUniqueInput{ID: p.Str(stages[0].ID)},
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	return db.client.Project(p.ProjectWhereUniqueInput{ID: p.Str(projectID)}).Exec(context.Background())
}

func (db *StateManager) AddStageToStage(projectID, stageID string, newStage tisp.NewStage) (*p.Project, error) {
	stages, err := db.client.Project(
		p.ProjectWhereUniqueInput{
			ID: p.Str(projectID),
		}).Stages(&p.StagesParamsExec{
		Where: &p.StageWhereInput{ID: p.Str(stageID)},
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	if len(stages) == 0 {
		return nil, errors.New("your stageID not exists inside of your projectID")
	}

	stage, err := db.client.CreateStage(p.StageCreateInput{
		Name:     newStage.Name,
		Selector: newStage.Selector,
	}).Exec(context.Background())
	if err != nil {
		return nil, err
	}

	if newStage.Tasks != nil {
		for _, t := range *newStage.Tasks {
			if _, err := db.AddTaskToStage(projectID, stage.ID, t); err != nil {
				return nil, err
			}
		}
	}

	_, err = db.client.UpdateStage(p.StageUpdateParams{
		Data: p.StageUpdateInput{
			Stages: &p.StageUpdateManyWithoutParentStageInput{
				Connect: []p.StageWhereUniqueInput{
					{ID: p.Str(stage.ID)},
				},
			},
		},
		Where: p.StageWhereUniqueInput{ID: p.Str(stages[0].ID)},
	}).Exec(context.Background())

	return db.client.Project(p.ProjectWhereUniqueInput{ID: p.Str(projectID)}).Exec(context.Background())
}
