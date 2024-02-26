package controllers

import (
	"encoding/json"
	"net/http"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/repo"
)

type BookController struct {
	bookRepo repo.IBook
}

func NewBookController(bookRepo repo.IBook) *BookController {
	return &BookController{bookRepo: bookRepo}
}

func (c *BookController) List(w http.ResponseWriter, r *http.Request) {
	books, _, err := c.bookRepo.Find(r.Context(), repo.PageOptions{})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
