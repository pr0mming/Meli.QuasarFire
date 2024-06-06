package services

import (
	"errors"
	"fmt"
	"math"
	"sort"
	"strings"

	"meli.quasarfire/src/aggregate"
	"meli.quasarfire/src/constants"
)

func (s *TopSecretService) GetEnemyStarship(parameters *aggregate.TopSecretRequest) (*aggregate.TopSecretResponse, error) {

	var (
		messages        [][]string
		distances       []float32
		starShipsPoints []map[string]float64
	)

	// Check: Is necessary least 2 points for get the location
	if len(parameters.Satellites) <= 1 {
		return nil, errors.New("sorry, we need at least two (2) points for get the location")
	}

	// Iterate satellites for extract message and distances
	for _, m := range parameters.Satellites {

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
	xDiff := math.Pow(float64(starShipsPoints[1]["x"]-starShipsPoints[0]["x"]), 2)
	yDiff := math.Pow(float64(starShipsPoints[1]["y"]-starShipsPoints[0]["y"]), 2)
	distanceStarShips := float32(math.Sqrt(xDiff + yDiff))

	// Append the distance in the first position,
	// because the first one is the distance between the starships the others is the user distances given
	distances = append([]float32{distanceStarShips}, distances...)

	x, y := s.GetLocation(GetLocationStarshipPoints{
		Point1: starShipsPoints[0],
		Point2: starShipsPoints[1],
	}, distances...)

	if math.IsNaN(float64(x)) || math.IsNaN(float64(y)) {
		return nil, errors.New("sorry, we cannot  calculate the position because any of the locations coords in NaN value")
	}

	decodedMessage := s.GetMessage(messages...)

	return &aggregate.TopSecretResponse{
		Position: aggregate.TopSecretCoordsResponse{
			X: x,
			Y: y,
		},
		Message: decodedMessage,
	}, nil

}

func (s *TopSecretService) GetLocation(points GetLocationStarshipPoints, distances ...float32) (x, y float32) {

	var (
		d1 float64 = float64(distances[0])
		d2 float64 = float64(distances[1])
		d3 float64 = float64(distances[2])

		point1 = points.Point1
		point2 = points.Point2
	)

	// Find cosin value using the Cosin Theorem ecuation
	cos := (math.Pow(d1, 2) + math.Pow(d2, 2) - math.Pow(d3, 2)) / (2 * d1 * d2)

	// With the cosin value we need calculate the Sin, we get two values (+ and -)
	sin := math.Sqrt(1 - math.Pow(cos, 2))

	// Use the factor (d2 / d1) and multiply by the sin components (+ and -)
	// Later plus this value with the first point
	xPoint := math.Round(point1["x"] + ((d2 / d1) * (cos*(point2["x"]-point1["x"]) - sin*(point2["y"]-point1["y"]))))
	yPoint := math.Round(point1["y"] + ((d2 / d1) * (sin*(point2["x"]-point1["x"]) + cos*(point2["y"]-point1["y"]))))

	return float32(xPoint), float32(yPoint)

}

func (s *TopSecretService) GetMessage(messages ...[]string) (msg string) {

	var (
		records       = make(map[int]map[string]int)
		decodedMesage []string
	)

	for _, m := range messages {

		for i, word := range m {

			wordTrimmed := strings.TrimSpace(word)

			if wordTrimmed != "" {

				message := records[i]

				// Save how many is repeated the word in the position
				if message != nil {

					message[wordTrimmed] += 1

				} else {

					message = map[string]int{wordTrimmed: 1}

				}

				records[i] = message

			}

		}

	}

	// Sort map by keys

	keys := make([]int, 0, len(records))

	for k := range records {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	lastWord := "" // Current word and used for avoid duplicates words next to other word

	for _, k := range keys {

		wordWeight := -1 // Check which word is most repeated in the position

		for k_k, v_v := range records[k] {

			if v_v > wordWeight && k_k != lastWord {
				wordWeight = v_v
				lastWord = k_k
			}

		}

		decodedMesage = append(decodedMesage, lastWord)

	}

	return strings.Join(decodedMesage, " ")

}
