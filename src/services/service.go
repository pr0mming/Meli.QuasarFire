package services

import "math"

func (s *TopSecretService) GetLocation(point1 map[string]float64, point2 map[string]float64, distances ...float32) (x, y float32) {

	var (
		d1 float64 = float64(distances[0])
		d2 float64 = float64(distances[1])
		d3 float64 = float64(distances[2])
	)

	cos := (math.Pow(d1, 2) + math.Pow(d2, 2) - math.Pow(d3, 2)) / (2 * d1 * d2)

	sin := math.Sqrt(1 - math.Pow(cos, 2))

	xPoint := math.Round(point1["x"] + ((d2 / d1) * (cos*(point2["x"]-point1["x"]) - sin*(point2["y"]-point1["y"]))))
	yPoint := math.Round(point1["y"] + ((d2 / d1) * (sin*(point2["x"]-point1["x"]) + cos*(point2["y"]-point1["y"]))))
	// xPoint := ((d2 / d1) * (sin*(point2["x"]-point1["x"]) + cos*(point2["y"]-point1["y"])))
	// yPoint := ((d2 / d1) * (cos*(point2["x"]-point1["x"]) - sin*(point2["y"]-point1["y"])))

	return float32(xPoint), float32(yPoint)

}

func (s *TopSecretService) GetMessage(messages ...[]string) (msg string) {

	return "Holo"

}
