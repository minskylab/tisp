package tisp

type ResourceType string

const Developer ResourceType = "developer"
const Scholarship ResourceType = "scholarship"
const ScrumMaster ResourceType = "scrum-master"

type Resource struct {
	ID         ID
	Selector   string // Selector acts like a unique index, is very useful
	Name       string
	MainType   ResourceType
	Types      []ResourceType
	Alias      string         // Alias is only extra information
	Experience ExperienceType //
	Cost       Cost
}
