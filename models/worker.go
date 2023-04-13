package models

type WorkerResponse struct {
	TaskID int     `json:"taskID"`
	Data   []byte  `json:"data"`
	Tasks  []*Task `json:"tasks"`
}

type Worker interface {
	CanExecute(task *Task) bool
	Execute(task *Task) *WorkerResponse
}
