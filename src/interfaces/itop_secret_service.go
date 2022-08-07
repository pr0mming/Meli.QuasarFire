package interfaces

import (
	"meli.quasarfire/src/aggregate"
	"meli.quasarfire/src/services"
)

type ITopSecretService interface {
	GetEnemyStarship(parameters *aggregate.TopSecretRequest) (*aggregate.TopSecretResponse, error)
	GetLocation(points services.GetLocationStarshipPoints, distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
