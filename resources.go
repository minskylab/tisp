package tisp

// type ResourceType string
//
// const Developer ResourceType = "developer"
// const Scholarship ResourceType = "scholarship"
// const ScrumMaster ResourceType = "scrum-master"
//
// type Resource struct {
// 	ID                       ID     `gorm:"primary_key"`
// 	Selector                 *string `gorm:"primary_key"` // Selector acts like a unique index, is very useful
// 	CreatedAt                time.Time
// 	UpdatedAt                time.Time
// 	Name                     string
// 	MainType                 ResourceType
// 	Types                    []ResourceType `gorm:"type:string[]"`
// 	Alias                    string         // Alias is only extra information
// 	Experience               ExperienceType //
// 	Cost                     Cost
// 	WorkingOnAsProjectLeader []Project `gorm:"foreignkey:LeaderID"`
// 	WorkingOnAsTaskLeader    []Task    `gorm:"many2many:task_resources;"`
// 	WorkingOn                []Task    `gorm:"foreignkey:LeaderID"`
// }
