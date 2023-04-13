package plugins

type Plugin interface {
	Execute(data interface{}) ([]byte, error)
}
