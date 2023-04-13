package models

type SampleAIWorker struct {
	Tags []string
}

func (w *SampleAIWorker) CanExecute(task *Task) bool {
	for _, taskTag := range task.Tags {
		for _, workerTag := range w.Tags {
			if taskTag == workerTag {
				return true
			}
		}
	}
	return false
}

func (w *SampleAIWorker) Execute(task *Task) *WorkerResponse {
	// Placeholder function to simulate AI agent task execution
	return &WorkerResponse{TaskID: task.ID, Data: nil}
}
