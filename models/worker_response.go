package models

type WorkerResponse struct {
	TaskID   uint     `json:"taskID"`
	WorkerID string   `json:"workerID"`
	Data     []string `json:"data"`
	Tasks    []*Task  `json:"Tasks"`
	Error    error    `json:"error"`
}

func (wr *WorkerResponse) IsError() bool {
	return wr.Error != nil
}
