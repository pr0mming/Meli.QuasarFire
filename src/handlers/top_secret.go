package handlers

import (
	"errors"
	"fmt"
	"math"
	"strings"

	"github.com/gin-gonic/gin"
	"meli.quasarfire/src/constants"
	"meli.quasarfire/src/handlers/aggregate"
	"meli.quasarfire/src/services"
)

// Handle for TopSecret godoc
// @Summary get the location and the decoded message
// @Description get the location and the decoded message using the right POST data
// @Tags Meli.QuasarFire
// @Accept json
// @Produce json
// @Param satellites body aggregate.TopSecretRequest true "Satellites JSON"
// @Success 200 {object} aggregate.TopSecretResponse
// @Router /topsecret [post]
func (s *TopSecretHandler) Handle(c *gin.Context) (interface{}, error) {

	var (
		requestBody     aggregate.TopSecretRequest
		messages        [][]string
		distances       []float32
		starShipsPoints []map[string]float64
	)

	// Check: Decode the request HTTP body
	err := c.BindJSON(&requestBody)
	if err != nil {
		return nil, fmt.Errorf("problems when try to deserialize the request body: %s", err.Error())
	}

	// Check: Is necessary least 2 points for get the location
	if len(requestBody.Satellites) <= 1 {
		return nil, errors.New("sorry, we need at least two (2) points for get the location")
	}

	// Iterate satellites for extract message and distances
	for _, m := range requestBody.Satellites {

		if len(m.Message) > 0 {
			messages = append(messages, m.Message)
		}

		point, ok := constants.STARSHIPS[strings.ToLower(m.Name)]

		if !ok {
			return nil, fmt.Errorf("sorry, but the ship name '%s' is unknown", m.Name)
		}

		// Check: The distances values must be great than 0
		if m.Distance <= 0 {
			return nil, fmt.Errorf("sorry, but the distance for the satellite %s must be great than 0", m.Name)
		}

		if len(starShipsPoints) < 2 {
			starShipsPoints = append(starShipsPoints, map[string]float64{"x": float64(point[0]), "y": float64(point[1])})
			distances = append(distances, m.Distance)

			continue
		}

		break

	}

	if len(messages) == 0 {
		return nil, fmt.Errorf("sorry, but we need at least one (1) item in the message for at least one (1) satellite for decode the message")
	}

	// Get the distance between the two starships (positions)
	_xDiff := math.Pow(float64(starShipsPoints[1]["x"]-starShipsPoints[0]["x"]), 2)
	_yDiff := math.Pow(float64(starShipsPoints[1]["y"]-starShipsPoints[0]["y"]), 2)
	_distanceStarShips := float32(math.Sqrt(_xDiff + _yDiff))

	// Append the distance in the first position,
	// because the first one is the distance between the starships the others is the user distances given
	distances = append([]float32{_distanceStarShips}, distances...)

	x, y := s.TopSecretService.GetLocation(services.GetLocationStarshipPoints{
		Point1: starShipsPoints[0],
		Point2: starShipsPoints[1],
	}, distances...)

	if math.IsNaN(float64(x)) || math.IsNaN(float64(y)) {
		return nil, errors.New("sorry, we cannot  calculate the position because any of the locations coords in NaN value")
	}

	decodedMessage := s.TopSecretService.GetMessage(messages...)

	return &aggregate.TopSecretResponse{
		Position: aggregate.TopSecretCoordsResponse{
			X: x,
			Y: y,
		},
		Message: decodedMessage,
	}, nil

}
