package tisp

import p "github.com/minskylab/tisp/dbclient"

type NewProjectInformation struct {
	Name      string
	Selector  *string
	PartnerID string
}

type NewResourceInformation struct {
	Name       string
	Cost       Cost
	MainType   p.ResourceType
	Experience float64
	Selector   *string
	Alias      *string
	Types      *[]p.ResourceType
}

type NewTaskInformation struct {
	Title       string
	Description string
	Selector    *string
	Resources   *[]string
	Leader      *string
	Children    *[]NewTaskInformation
}


type NewPartnerInformation struct {
	CompanyName string
	Selector *string

	Email *string
	Phone *string
}