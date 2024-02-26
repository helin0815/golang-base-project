package repo

import (
	"context"

	"gorm.io/gen"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal/query"
)

type PageOptions struct {
	PageNo   int
	PageSize int
}

func (p PageOptions) PageLimit() int {
	return p.PageSize
}

func (p PageOptions) PageOffset() int {
	return (p.PageNo - 1) * p.PageSize
}

type Option interface {
	buildConds(q *query.Query) []gen.Condition
}

type IDType interface {
	int | int64 | string
}

type CURD[T comparable, ID IDType, O any] interface {
	Writer[T, ID]
	Reader[T, ID, O]
	Deleter[ID]
}

type Writer[T comparable, ID IDType] interface {
	Creator[T]
	Updater[T, ID]
	Deleter[ID]
}

type Reader[T comparable, ID IDType, O any] interface {
	Getter[T, ID]
	Finder[T, O]
}

type Getter[T comparable, ID IDType] interface {
	Get(ctx context.Context, id ID) (T, error)
}

type Finder[T comparable, O any] interface {
	Find(ctx context.Context, opts O) ([]T, int64, error)
}

type Creator[T comparable] interface {
	Create(ctx context.Context, entity T) error
}

type Updater[T comparable, ID IDType] interface {
	Update(ctx context.Context, entity T) error
}

type Deleter[ID IDType] interface {
	Delete(ctx context.Context, id ID) error
}

type FakeRepo[T comparable, ID IDType, O any] struct {
	CreateFn func(ctx context.Context, entity T) error
	UpdateFn func(ctx context.Context, entity T) error
	DeleteFn func(ctx context.Context, id ID) error
	GetFn    func(ctx context.Context, id ID) (T, error)
	FindFn   func(ctx context.Context, opts O) ([]T, int64, error)
}

func (f FakeRepo[T, ID, O]) Create(ctx context.Context, entity T) error {
	return f.CreateFn(ctx, entity)
}

func (f FakeRepo[T, ID, O]) Update(ctx context.Context, entity T) error {
	return f.UpdateFn(ctx, entity)
}

func (f FakeRepo[T, ID, O]) Delete(ctx context.Context, id ID) error {
	return f.DeleteFn(ctx, id)
}

func (f FakeRepo[T, ID, O]) Get(ctx context.Context, id ID) (T, error) {
	return f.GetFn(ctx, id)
}

func (f FakeRepo[T, ID, O]) Find(ctx context.Context, opts O) ([]T, int64, error) {
	return f.FindFn(ctx, opts)
}
