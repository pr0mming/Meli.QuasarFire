package provider

import "meli.quasarfire/src/interfaces"

type Server struct {
	TopSecretHandler      interfaces.ITopSecretHandler
	TopSecretSplitHandler interfaces.ITopSecretSplitHandler
}
