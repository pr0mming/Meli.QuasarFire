//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package injector

import (
	"github.com/google/wire"
	"meli.quasarfire/src/handlers"
	"meli.quasarfire/src/interfaces"
	"meli.quasarfire/src/provider"
	"meli.quasarfire/src/services"
)

func Wire() *provider.Server {
	panic(wire.Build(
		wire.NewSet(
			provider.ProvideTopSecretHandler,
			provider.ProvideTopSecretSplitHandler,
			provider.ProvideService,
			provider.ProvideServer,

			wire.Bind(new(interfaces.ITopSecretHandler), new(*handlers.TopSecretHandler)),
			wire.Bind(new(interfaces.ITopSecretSplitHandler), new(*handlers.TopSecretSplitHandler)),
			wire.Bind(new(interfaces.ITopSecretService), new(*services.TopSecretService)),
			wire.Bind(new(interfaces.IServer), new(*provider.Server)),
		),
	))
}
