//go:build wireinject
// +build wireinject

package app

import (
	"context"

	"github.com/google/wire"

	"gitlabee.chehejia.com/k8s/liks-gitops/internal/config"
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/controllers"
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/dal"
	"gitlabee.chehejia.com/k8s/liks-gitops/internal/repo"
)

func InitializeServer(ctx context.Context) (*Server, error) {
	wire.Build(
		config.New,
		dal.NewQuery,
		repo.NewRegistry,
		controllers.NewRegistry,
		controllers.New,
		NewServer,
	)

	return &Server{}, nil
}
