package models

type Board struct {
	Tasks      []*Task
	currentIdx int
}

//Add TODO Sort by priority
func (b *Board) Add(t *Task) {
	b.Tasks = append(b.Tasks, t)
}

//Map returns the board map representation of the Tasks
func (b *Board) Map() map[Status][]*Task {
	m := map[Status][]*Task{}

	for _, t := range b.Tasks {
		if _, ok := m[t.Status]; !ok {
			m[t.Status] = []*Task{t}
			continue
		}
		m[t.Status] = append(m[t.Status], t)
	}

	return m
}

func (b *Board) Backlog() []*Task {
	return b.ofStatus(Backlog)
}

func (b *Board) InProgress() []*Task {
	return b.ofStatus(InProgress)
}

func (b *Board) Testing() []*Task {
	return b.ofStatus(Testing)
}

func (b *Board) Completed() []*Task {
	return b.ofStatus(Completed)
}

func (b *Board) ofStatus(status Status) []*Task {
	tasks := []*Task{}
	for _, t := range b.Tasks {
		if t.Status == status {
			tasks = append(tasks, t)
		}
	}

	return tasks
}
