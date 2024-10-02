package controller

import "github.com/VuKhoa23/advanced-web-be/internal/controller/http"

type ApiContainer struct {
	HttpServer *http.Server
}

func NewApiContainer(httpServer *http.Server) *ApiContainer {
	return &ApiContainer{HttpServer: httpServer}
}
