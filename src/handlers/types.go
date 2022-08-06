package handlers

import (
	interfaces "meli.quasarfire/src/interfaces"
)

type TopSecretSatelliteRequest struct {
	Name     string
	Distance float32
	Message  []string
}

type TopSecretRequest struct {
	Satellites []TopSecretSatelliteRequest
}

type TopSecretCoordsResponse struct {
	X float32
	Y float32
}

type TopSecretResponse struct {
	Position TopSecretCoordsResponse
	Message  string
}

type TopSecretHandler struct {
	TopSecretService interfaces.ITopSecretService
}
