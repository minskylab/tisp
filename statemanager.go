package tisp

import p "github.com/minskylab/tisp/repository"

type StateManager interface {
	RegisterNewProject(info NewProjectInformation) (*p.Project, error)
	RegisterNewResource(newRes NewResourceInformation) (*p.Resource, error)
	RegisterNewPartner(newClient NewPartnerInformation) (*p.Partner, error)

	AddTaskToProject(projectID string, task NewTaskInformation) (*p.Project, error)
	AddSubTaskToTask(taskID string, task NewTaskInformation) (*p.Task, error)
	AddResourceToTask(taskID string, res NewResourceInformation) (*p.Task, error)
	// ++
	AddStageToProject(projectID string, newStage NewStage) (*p.Project, error)
	AddTaskToStage(projectID, stageID string, task NewTaskInformation) (*p.Project, error)
	AddStageToStage(projectID, stageID string, newStage NewStage) (*p.Project, error)

	GetProjects(p ...Selector) ([]p.Project, error)
	GetProject(projectID string) (*p.Project, error)
	GetProjectBySelector(selector string) (*p.Project, error)

	GetResources(p ...Selector) ([]p.Resource, error)
	GetResource(resID string) (*p.Resource, error)
	GetResourceBySelector(selector string) (*p.Resource, error)

	GetTasks(p ...Selector) ([]p.Task, error)
	GetTask(taskID string) (*p.Task, error)
	GetTaskBySelector(selector string) (*p.Task, error)

	GetPartners(p ...Selector) ([]p.Partner, error)
	GetPartner(partnerID string) (*p.Partner, error)
	GetPartnerBySelector(selector string) (*p.Partner, error)


}

