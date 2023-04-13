package models

type Task struct {
	ID       int      `json:"ID"`
	Name     string   `json:"name"`
	Tags     []string `json:"tags"`
	Inputs   map[Action]interface{}
	Outputs  map[Action]interface{}
	Priority uint    `json:"priority"`
	Children []*Task `json:"children"`
}

func (t *Task) Complete() bool {
	for _, st := range t.Children {
		if !st.Complete() {
			return false
		}
	}
	return true
}
