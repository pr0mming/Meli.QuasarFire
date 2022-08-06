package handlers

import (
	"errors"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
)

var ships = map[string][]float32{
	"kenobi":    {-5, 4},
	"skywalker": {7, 3},
	"sato":      {500, 100},
}

func (s *TopSecretHandler) Handle(c *gin.Context) (interface{}, error) {

	var (
		requestBody TopSecretRequest
		messages    [][]string
		distances   []float32
		shipsPoints []map[string]float64
	)

	// Decode the request HTTP body
	err := c.BindJSON(&requestBody)
	if err != nil {
		return nil, err
	}

	// Check: Is necessary least 2 points for get the location
	if len(requestBody.Satellites) <= 1 {
		return nil, errors.New("cannot decode the message")
	}

	// Iterate satellites for extract message and distances
	for _, m := range requestBody.Satellites {
		messages = append(messages, m.Message)

		point, ok := ships[strings.ToLower(m.Name)]

		if !ok {
			return nil, errors.New("unknow ship")
		}

		if len(shipsPoints) < 2 {
			shipsPoints = append(shipsPoints, map[string]float64{"x": float64(point[0]), "y": float64(point[1])})
			distances = append(distances, m.Distance)

			continue
		}

		break

	}

	_xDiff := math.Pow(float64(shipsPoints[1]["x"]-shipsPoints[0]["x"]), 2)
	_yDiff := math.Pow(float64(shipsPoints[1]["y"]-shipsPoints[0]["y"]), 2)
	_distanceShips := float32(math.Sqrt(_xDiff + _yDiff))

	distances = append([]float32{_distanceShips}, distances...)

	x, y := s.TopSecretService.GetLocation(shipsPoints[0], shipsPoints[1], distances...)
	decodedMessage := s.TopSecretService.GetMessage(messages...)

	return &TopSecretResponse{
		Position: TopSecretCoordsResponse{
			X: x,
			Y: y,
		},
		Message: decodedMessage,
	}, nil

}
