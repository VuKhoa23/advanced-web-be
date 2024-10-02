//go:build wireinject
// +build wireinject

package internal

import (
	"github.com/VuKhoa23/advanced-web-be/internal/controller"
	"github.com/VuKhoa23/advanced-web-be/internal/controller/http"
	v1 "github.com/VuKhoa23/advanced-web-be/internal/controller/http/v1"
	"github.com/VuKhoa23/advanced-web-be/internal/database"
	repositoryimplement "github.com/VuKhoa23/advanced-web-be/internal/repository/implement"
	serviceimplement "github.com/VuKhoa23/advanced-web-be/internal/service/implement"
	"github.com/google/wire"
)

var container = wire.NewSet(
	controller.NewApiContainer,
)

var controllerSet = wire.NewSet(
	http.NewServer,
)

var handlerSet = wire.NewSet(
	v1.NewStudentHandler,
)

var serviceSet = wire.NewSet(
	serviceimplement.NewStudentService,
)

var repositorySet = wire.NewSet(
	repositoryimplement.NewStudentRepository,
)

func InitializeContainer(
	db database.Db,
) *controller.ApiContainer {
	wire.Build(controllerSet, handlerSet, serviceSet, repositorySet, container)
	return &controller.ApiContainer{}
}
