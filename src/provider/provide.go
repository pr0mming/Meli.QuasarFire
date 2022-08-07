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

	hdlTopSecret          *handlers.TopSecretHandler
	hdlTopSecretOnce      sync.Once
	hdlTopSecretSplit     *handlers.TopSecretSplitHandler
	hdlTopSecretSplitOnce sync.Once

	topSecretService     *services.TopSecretService
	topSecretServiceOnce sync.Once
)

// Handlers

func ProvideTopSecretHandler(s interfaces.ITopSecretService) *handlers.TopSecretHandler {
	hdlTopSecretOnce.Do(func() {
		hdlTopSecret = &handlers.TopSecretHandler{
			TopSecretService: s,
		}
	})

	return hdlTopSecret
}

func ProvideTopSecretSplitHandler(s interfaces.ITopSecretService) *handlers.TopSecretSplitHandler {
	hdlTopSecretSplitOnce.Do(func() {
		hdlTopSecretSplit = &handlers.TopSecretSplitHandler{
			TopSecretService: s,
		}
	})

	return hdlTopSecretSplit
}

// Services and others

func ProvideService() *services.TopSecretService {
	topSecretServiceOnce.Do(func() {
		topSecretService = &services.TopSecretService{}
	})

	return topSecretService
}

func ProvideServer(
	_hdlTopSecret interfaces.ITopSecretHandler,
	_hdlTopSecretSplit interfaces.ITopSecretSplitHandler,
) *Server {
	serverOnce.Do(func() {
		server = &Server{
			TopSecretHandler:      _hdlTopSecret,
			TopSecretSplitHandler: _hdlTopSecretSplit,
		}
	})

	return server
}
