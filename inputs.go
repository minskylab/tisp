package tisp

import p "github.com/minskylab/tisp/repository"

type NewProjectInformation struct {
	Name      string  `json:"name" yaml:"name"`
	Selector  *string `json:"selector" yaml:"selector"`
	PartnerID string  `json:"partner" yaml:"partner"`
}

type NewResourceInformation struct {
	Name       string            `json:"name" yaml:"name"`
	Cost       Cost              `json:"cost" yaml:"cost"`
	MainType   p.ResourceType    `json:"mainType" yaml:"mainType"`
	Experience float64           `json:"experience" yaml:"experience"`
	Selector   *string           `json:"selector" yaml:"selector"`
	Alias      *string           `json:"alias" yaml:"alias"`
	Types      *[]p.ResourceType `json:"types" yaml:"types"`
}

type NewTaskInformation struct {
	Title       string                `json:"title" yaml:"title"`
	Description string                `json:"description" yaml:"description"`
	Selector    *string               `json:"selector" yaml:"selector"`
	Resources   *[]string             `json:"resources" yaml:"resources"`
	Leader      *string               `json:"leader" yaml:"leader"`
	Children    *[]NewTaskInformation `json:"tasks" yaml:"tasks"`
}

type NewPartnerInformation struct {
	CompanyName string  `json:"name" yaml:"name"`
	Selector    *string `json:"selector" yaml:"selector"`

	Email *string `json:"email" yaml:"email"`
	Phone *string `json:"phone" yaml:"phone"`
}

type NewStage struct {
	Name     string                `json:"name" yaml:"name"`
	Selector *string               `json:"selector" yaml:"selector"`
	Tasks    *[]NewTaskInformation `json:"tasks" yaml:"tasks"`
}
