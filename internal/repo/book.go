package repo

import (
	"context"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal/query"
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/entities"
)

type IBook interface {
	Reader[*entities.Book, int64, PageOptions]
}

var _ IBook = (*Book)(nil)

type Book struct {
	q *query.Query
}

func NewBook(q *query.Query) *Book {
	return &Book{q: q}
}

func (b *Book) Get(ctx context.Context, id int64) (*entities.Book, error) {
	return b.q.Book.WithContext(ctx).Where(b.q.Book.Id.Eq(id)).First()
}

func (b *Book) Find(ctx context.Context, opts PageOptions) ([]*entities.Book, int64, error) {
	return b.q.Book.WithContext(ctx).FindByPage(opts.PageOffset(), opts.PageLimit())
}
