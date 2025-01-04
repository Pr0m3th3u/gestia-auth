package handlers

type IRootHandler interface {
}

type RootHandler struct {
}

var (
	_ IRootHandler = (*RootHandler)(nil)
)

func NewRootHandler() IRootHandler {
	return &RootHandler{}
}
