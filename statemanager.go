package tisp

import p "github.com/minskylab/tisp/dbclient"

type StateManager interface {
	RegisterNewProject(info NewProjectInformation) (*p.Project, error)
	RegisterNewResource(newRes NewResourceInformation) (*p.Resource, error)
	AddTaskToProject(projectID string, task NewTaskInformation) (*p.Project, error)
	AddSubTaskToTask(taskID string, task NewTaskInformation) (*p.Task, error)
	AddResourceToTask(taskID string, res NewResourceInformation) (*p.Task, error)

	GetProjects(p ...Pagination) ([]p.Project, error)
	GetProject(projectID string) (*p.Project, error)
	GetProjectBySelector(selector string) (*p.Project, error)

	GetResources(p ...Pagination) ([]p.Resource, error)
	GetResource(resID string) (*p.Resource, error)
	GetResourceBySelector(selector string) (*p.Resource, error)

	GetTasks(p ...Pagination) ([]p.Task, error)
	GetTask(taskID string) (*p.Task, error)
	GetTaskBySelector(selector string) (*p.Task, error)
}
