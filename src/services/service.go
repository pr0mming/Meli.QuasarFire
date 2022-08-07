package services

import (
	"math"
	"strings"
)

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

			_word := strings.TrimSpace(word)

			if _word != "" {

				message := records[i]

				if message != nil {

					message[_word] += 1

				} else {

					message = map[string]int{_word: 1}

				}

				records[i] = message

			}

		}

	}

	for _, v := range records {

		wordWeight := -1
		word := ""
		lastWord := ""

		for k_k, v_v := range v {

			if v_v > wordWeight && k_k != lastWord {
				wordWeight = v_v
				word = k_k
				lastWord = k_k
			}

		}

		decodedMesage = append(decodedMesage, word)

	}

	return strings.Join(decodedMesage, " ")

}
