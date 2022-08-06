package interfaces

type ITopSecretService interface {
	GetLocation(point1 map[string]float64, point2 map[string]float64, distances ...float32) (x, y float32)
	GetMessage(messages ...[]string) (msg string)
}
