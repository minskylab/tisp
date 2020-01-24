package tisp

import p "github.com/minskylab/tisp/dbclient"

type DB interface {
	RegisterNewProject(info NewProjectInformation) (*p.Project, error)
	RegisterNewResource(newRes NewResourceInformation) (*p.Resource, error)
	AddTaskToProject(projectID ID, task NewTaskInformation) (*p.Project, error)
	AddSubTaskToTask(taskID ID, task NewTaskInformation) (*p.Task, error)
	AddResourceToTask(taskID ID, res NewResourceInformation) (*p.Task, error)

	GetProject(projectID ID) (*p.Project, error)
	GetProjectBySelector(selector string) (*p.Project, error)

	GetResource(resID ID) (*p.Resource, error)
	GetResourceBySelector(selector string) (*p.Resource, error)

	GetTask(taskID ID) (*p.Task, error)
	GetTaskBySelector(selector string) (*p.Task, error)
}
