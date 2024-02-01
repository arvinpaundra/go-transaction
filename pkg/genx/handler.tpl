package {{.PackageName}}


import (
	"{{.ModuleName}}/internal/factory"
)

type handler struct {
	service Service
}

func NewHandler(f *factory.Factory) *handler {
	return &handler{
		service: NewService(f),
	}
}
