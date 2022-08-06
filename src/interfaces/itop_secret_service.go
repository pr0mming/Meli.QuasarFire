package interfaces

import "meli.quasarfire/src/services"

type ITopSecretService interface {
	GetLocation(points services.GetLocationStarshipPoints, distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
