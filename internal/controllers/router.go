package controllers

import (
	"net/http"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/repo"
)

type Registry struct {
	book *BookController
}

func NewRegistry(repoRegistry *repo.Registry) *Registry {
	return &Registry{
		book: NewBookController(repoRegistry.Book),
	}
}

func New(reg *Registry) http.Handler {
	// 默认为官方的路由选择器，您可以在这里换成任何您想要的路由，如 gin,echo 等
	mux := http.NewServeMux()
	mux.HandleFunc("/", reg.book.List)
	return mux
}
