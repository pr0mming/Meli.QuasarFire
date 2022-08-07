package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"meli.quasarfire/src/aggregate"
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
		requestBody aggregate.TopSecretRequest
	)

	// Check: Decode the request HTTP body
	err := c.BindJSON(&requestBody)
	if err != nil {
		return nil, fmt.Errorf("problems when try to deserialize the request body: %s", err.Error())
	}

	answer, err := s.TopSecretService.GetEnemyStarship(&requestBody)
	if err != nil {
		return nil, err
	}

	return answer, nil

}
