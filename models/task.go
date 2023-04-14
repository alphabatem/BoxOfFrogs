package models

import "log"

type Task struct {
	ID       uint     `json:"ID"`
	Tags     []string `json:"tags"`
	Input    string   `json:"input"`
	Priority uint     `json:"priority"`
	Children []*Task  `json:"children"`

	Status Status `json:"status"`
}

func NewTask(prompt string, tags []string, id uint) *Task {

	log.Printf("#%v - New Task: %v", id, tags)

	return &Task{
		ID:       id,
		Tags:     tags,
		Input:    prompt,
		Children: []*Task{},
		Status:   Backlog,
	}
}

func (t *Task) Complete() bool {
	for _, st := range t.Children {
		if !st.Complete() {
			return false
		}
	}
	return true
}
