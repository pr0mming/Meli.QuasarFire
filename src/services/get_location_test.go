package services

import (
	"fmt"
	"testing"

	"meli.quasarfire/src/constants"
)

type GetLocationTestOutput struct {
	x float32
	y float32
}

type GetLocationTestInput struct {
	arg1     GetLocationStarshipPoints
	arg2     float32
	arg3     float32
	arg4     float32
	expected GetLocationTestOutput
}

func TestGetLocation(t *testing.T) {

	var (
		kenobi_coords    = constants.STARSHIPS["kenobi"]
		skywalker_coords = constants.STARSHIPS["skywalker"]

		starships_coords = GetLocationStarshipPoints{
			Point1: map[string]float64{"x": float64(kenobi_coords[0]), "y": float64(kenobi_coords[1])},
			Point2: map[string]float64{"x": float64(skywalker_coords[0]), "y": float64(skywalker_coords[1])},
		}

		distanceStarships float32 = 608.28
	)

	var testList = []GetLocationTestInput{
		{
			starships_coords,
			distanceStarships,
			910.14,
			389.30,
			GetLocationTestOutput{x: 300, y: 234},
		},
		{
			starships_coords,
			distanceStarships,
			997.01,
			571.70,
			GetLocationTestOutput{x: 256, y: 450},
		},
		{
			starships_coords,
			distanceStarships,
			768.37,
			514.20,
			GetLocationTestOutput{x: -20, y: 400},
		},
	}

	service := TopSecretService{}

	for i, test := range testList {

		fmt.Printf("[+] Testing case %v of %v \n", (i + 1), len(testList))

		x, y := service.GetLocation(test.arg1, test.arg2, test.arg3, test.arg4)

		if x != test.expected.x {
			t.Errorf("Output in X component %f not equals to expected %f!", x, test.expected.x)
		}

		if y != test.expected.y {
			t.Errorf("Output in Y component %f not equals to expected %f!", x, test.expected.y)
		}

	}

}
