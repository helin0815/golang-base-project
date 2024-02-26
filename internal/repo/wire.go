package repo

import (
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal/query"
)

type Registry struct {
	Book IBook
}

func NewRegistry(q *query.Query) *Registry {
	return &Registry{
		Book: NewBook(q),
	}
}
