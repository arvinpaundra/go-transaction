package {{.PackageName}}

import (
	"{{.ModuleName}}/internal/factory"
	"{{.ModuleName}}/internal/repository"
)

type service struct {
	{{.EntityName}}Repository repository.{{.EntityName}}
}

type Service interface {
}

// A function to call factory to initialize database connection to this/these repository
func NewService(f *factory.Factory) Service {
	return &service{
		{{.EntityName}}Repository: f.{{.EntityName}}Repository,
	}
}
