package handlers

import (
	interfaces "meli.quasarfire/src/interfaces"
)

type TopSecretHandler struct {
	TopSecretService interfaces.ITopSecretService
}

type TopSecretSplitHandler struct {
	TopSecretService interfaces.ITopSecretService
}
