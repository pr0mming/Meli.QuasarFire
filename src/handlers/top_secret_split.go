package handlers

import "github.com/gin-gonic/gin"

// Handle for TopSecretSplit godoc
// @Summary get the location and the decoded message
// @Description get the location and the decoded message using the right POST data
// @Tags Meli.QuasarFire
// @Accept json
// @Produce json
// @Param satellites body aggregate.TopSecretSplitSatelliteRequest true "Satellites JSON"
// @Success 200 {object} aggregate.TopSecretResponse
// @Router /topsecret_split [get]
func (s *TopSecretSplitHandler) Handle(c *gin.Context) (interface{}, error) {

	data := c.Param("fs_name")

	return data, nil

}
