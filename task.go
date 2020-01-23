package tisp

type TaskState string

type Task struct {
	ID          ID
	Title       string
	Selector    string // unique
	Description string

	ParentTask ID

	Leader    ID
	Resources map[string]string

	State    TaskState
	Children []ID
}
