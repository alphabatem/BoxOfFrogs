package workers

//GetWorker returns a registered worker
//When developing a worker make sure to register it here
func GetWorker(id string) Worker {
	switch id {
	case OpenAiWorker{}.Id():
		return &OpenAiWorker{}
	case TaskWorker{}.Id():
		return &TaskWorker{}
	case DeveloperWorker{}.Id():
		return &DeveloperWorker{}
	default:
		return nil
	}
}
