package services

import (
	"fmt"
	"testing"
)

type GetMessageTestInput struct {
	input    [][]string
	expected string
}

func TestGetMessage(t *testing.T) {

	var testList = []GetMessageTestInput{
		{
			input: [][]string{
				{
					"",
					"mensaje",
					"",
					"secreto",
				},
				{
					"es",
					"",
					"",
					"secreto",
				},
				{
					"",
					"mensaje",
					"",
					"secreto",
				},
			},
			expected: "es mensaje secreto",
		},
	}

	service := TopSecretService{}

	for i, test := range testList {

		fmt.Printf("[+] Testing case %v of %v \n", (i + 1), len(testList))

		message := service.GetMessage(test.input...)

		if message != test.expected {
			t.Errorf("Output '%s' not equals to expected '%s'!", message, test.expected)
		}

	}

}
