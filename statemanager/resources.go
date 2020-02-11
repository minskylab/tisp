package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"
)

func (db *StateManager) RegisterNewResource(newRes tisp.NewResourceInformation) (*p.Resource, error) {
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


func (db *StateManager) AddResourceToTask(taskID string, res tisp.NewResourceInformation) (*p.Task, error) {
	types := new(p.ResourceCreatetypesInput)
	if res.Types != nil {
		for _, t := range *res.Types {
			types.Set = append(types.Set, t)
		}
	}

	return db.client.UpdateTask(p.TaskUpdateParams{
		Where: p.TaskWhereUniqueInput{
			ID: p.Str(taskID),
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
