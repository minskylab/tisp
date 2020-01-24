package tisp

type TaskState string

// type Task struct {
// 	ID        ID `gorm:"primary_key"`
// 	CreatedAt time.Time
// 	UpdatedAt time.Time
//
// 	Title       string
// 	Selector    *string `gorm:"primary_key"` // unique
// 	Description string
//
// 	ParentTask ID
//
// 	LeaderID  ID
// 	Resources []Resource `gorm:"many2many:task_resources;"`
//
// 	State    TaskState
// 	Children []Task `gorm:"foreignkey:ParentTask"`
//
// 	ProjectParent ID
// }
