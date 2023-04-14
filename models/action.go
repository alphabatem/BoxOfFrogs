package models

type Action int

const (
	ActionDefault = iota
	Description   = iota
	FileSave
)

type Status string

const (
	Backlog    = "backlog"
	InProgress = "in-progress"
	Testing    = "testing"
	Completed  = "completed"
)
