package provider

import (
	"sync"

	"meli.quasarfire/src/handlers"
	"meli.quasarfire/src/interfaces"
	"meli.quasarfire/src/services"
)

var (
	server     *Server
	serverOnce sync.Once

	hdlTopSecret     *handlers.TopSecretHandler
	hdlTopSecretOnce sync.Once

	topSecretService     *services.TopSecretService
	topSecretServiceOnce sync.Once
)

func ProvideHandler(s interfaces.ITopSecretService) *handlers.TopSecretHandler {
	hdlTopSecretOnce.Do(func() {
		hdlTopSecret = &handlers.TopSecretHandler{
			TopSecretService: s,
		}
	})

	return hdlTopSecret
}

func ProvideService() *services.TopSecretService {
	topSecretServiceOnce.Do(func() {
		topSecretService = &services.TopSecretService{}
	})

	return topSecretService
}

func ProvideServer(hdl interfaces.ITopSecretHandler) *Server {
	serverOnce.Do(func() {
		server = &Server{
			TopSecretHandler: hdl,
		}
	})

	return server
}
