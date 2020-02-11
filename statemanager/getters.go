package statemanager

import (
	"context"

	"github.com/minskylab/tisp"
	p "github.com/minskylab/tisp/repository"

)

func (db *StateManager) GetProject(projectID string) (*p.Project, error) {
	return db.client.Project(p.ProjectWhereUniqueInput{
		ID: p.Str(projectID),
	}).Exec(context.Background())
}

func (db *StateManager) GetProjectBySelector(selector string) (*p.Project, error) {
	return db.client.Project(p.ProjectWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *StateManager) GetResource(resID string) (*p.Resource, error) {
	return db.client.Resource(p.ResourceWhereUniqueInput{
		ID: p.Str(resID),
	}).Exec(context.Background())
}

func (db *StateManager) GetResourceBySelector(selector string) (*p.Resource, error) {
	return db.client.Resource(p.ResourceWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *StateManager) GetTask(taskID string) (*p.Task, error) {
	return db.client.Task(p.TaskWhereUniqueInput{
		ID: p.Str(taskID),
	}).Exec(context.Background())
}

func (db *StateManager) GetTaskBySelector(selector string) (*p.Task, error) {
	return db.client.Task(p.TaskWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *StateManager) GetPartner(partnerID string) (*p.Partner, error) {
	return db.client.Partner(p.PartnerWhereUniqueInput{
		ID: p.Str(partnerID),
	}).Exec(context.Background())
}

func (db *StateManager) GetPartnerBySelector(selector string) (*p.Partner, error) {
	return db.client.Partner(p.PartnerWhereUniqueInput{
		Selector: p.Str(selector),
	}).Exec(context.Background())
}

func (db *StateManager) GetProjects(pag ...tisp.Pagination) ([]p.Project, error) {
	var pagination = tisp.DefaultPagination

	if len(pag) != 0 {
		pagination = pag[0] // TODO
	}

	var where *p.ProjectWhereInput

	if pagination.ByIDs != nil {
		where = &p.ProjectWhereInput{
			IDIn: *pagination.ByIDs,
		}
	}

	if pagination.BySelectors != nil {
		if where == nil {
			where = &p.ProjectWhereInput{}
		}
		where.SelectorIn = *pagination.BySelectors
	}

	return db.client.Projects(&p.ProjectsParams{
		Where:   where,
		Skip:    p.Int32(pagination.Skip),
		First:   p.Int32(pagination.First),
	}).Exec(context.Background())
}

func (db *StateManager) GetResources(pag ...tisp.Pagination) ([]p.Resource, error) {
	var pagination = tisp.DefaultPagination

	if len(pag) != 0 {
		pagination = pag[0] // TODO
	}

	var where *p.ResourceWhereInput

	if pagination.ByIDs != nil {
		where = &p.ResourceWhereInput{
			IDIn: *pagination.ByIDs,
		}
	}

	if pagination.BySelectors != nil {
		if where == nil {
			where = &p.ResourceWhereInput{}
		}
		where.SelectorIn = *pagination.BySelectors
	}

	if pagination.First == 0 {
		pagination.First = tisp.DefaultPagination.First
	}

	if pagination.Skip == 0 {
		pagination.Skip = tisp.DefaultPagination.Skip
	}

	return db.client.Resources(&p.ResourcesParams{
		Where:   where,
		Skip:    p.Int32(pagination.Skip),
		First:   p.Int32(pagination.First),
	}).Exec(context.Background())
}

func (db *StateManager) GetTasks(pag ...tisp.Pagination) ([]p.Task, error) {
	var pagination = tisp.DefaultPagination

	if len(pag) != 0 {
		pagination = pag[0] // TODO
	}

	var where *p.TaskWhereInput

	if pagination.ByIDs != nil {
		where = &p.TaskWhereInput{
			IDIn: *pagination.ByIDs,
		}
	}

	if pagination.BySelectors != nil {
		if where == nil {
			where = &p.TaskWhereInput{}
		}
		where.SelectorIn = *pagination.BySelectors
	}

	return db.client.Tasks(&p.TasksParams{
		Where:   where,
		Skip:    p.Int32(pagination.Skip),
		First:   p.Int32(pagination.First),
	}).Exec(context.Background())
}


func (db *StateManager) GetPartners(pag ...tisp.Pagination) ([]p.Partner, error) {
	var pagination = tisp.DefaultPagination

	if len(pag) != 0 {
		pagination = pag[0] // TODO
	}

	var where *p.PartnerWhereInput

	if pagination.ByIDs != nil {
		where = &p.PartnerWhereInput{
			IDIn: *pagination.ByIDs,
		}
	}

	if pagination.BySelectors != nil {
		if where == nil {
			where = &p.PartnerWhereInput{}
		}
		where.SelectorIn = *pagination.BySelectors
	}

	return db.client.Partners(&p.PartnersParams{
		Where:   where,
		Skip:    p.Int32(pagination.Skip),
		First:   p.Int32(pagination.First),
	}).Exec(context.Background())
}

