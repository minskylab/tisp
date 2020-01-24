package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/dbclient"
)

func (db *StateManager) RegisterNewProject(info tisp.NewProjectInformation) (*p.Project, error) {
	return db.client.CreateProject(p.ProjectCreateInput{
		Selector: info.Selector,
		Name:     info.Name,
		Client: p.CompanyCreateOneWithoutProjectsInput{
			Connect: &p.CompanyWhereUniqueInput{
				ID: p.Str(info.ClientID),
			},
		},
	}).Exec(context.Background())
}


