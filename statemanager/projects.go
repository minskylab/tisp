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
		Partner: p.PartnerCreateOneWithoutProjectsInput{
			Connect: &p.PartnerWhereUniqueInput{
				ID: p.Str(info.PartnerID),
			},
		},
	}).Exec(context.Background())
}


