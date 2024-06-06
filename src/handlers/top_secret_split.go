package handlers

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"meli.quasarfire/src/aggregate"
)

// Handle for TopSecretSplit godoc
// @Summary get the location and the decoded message
// @Description get the location and the decoded message using the right POST data
// @Tags Meli.QuasarFire
// @Produce json
// @Param fs_name query string true "First Satellite name"
// @Param fs_distance query number true "First Satellite distance"
// @Param fs_message query string true "First Satellite message"
// @Param ss_name query string true "Second Satellite name"
// @Param ss_distance query number true "Second Satellite distance"
// @Param ss_message query string true "Second Satellite message"
// @Param ts_name query string true "Third Satellite name"
// @Param ts_distance query number true "Third Satellite distance"
// @Param ts_message query string true "Third Satellite message"
// @Success 200 {object} aggregate.TopSecretResponse
// @Router /topsecret_split [get]
func (s *TopSecretSplitHandler) HandleGet(c *gin.Context) (interface{}, error) {

	var (
		satelletiesOrderKeys = []string{"fs", "ss", "ts"}

		data aggregate.TopSecretRequest
	)

	for _, v := range satelletiesOrderKeys {

		name := c.Query(fmt.Sprintf("%s_%s", v, "name"))

		if name != "" {

			distance, err := strconv.ParseFloat(c.Query(fmt.Sprintf("%s_%s", v, "distance")), 32)
			if err != nil {
				return nil, fmt.Errorf("problems when try to convert distance to number: %s", err.Error())
			}

			message := c.Query(fmt.Sprintf("%s_%s", v, "message"))

			messageArr := strings.Split(strings.ReplaceAll(message, "\"", ""), ",")

			data.Satellites = append(data.Satellites, aggregate.TopSecretSatelliteRequest{
				Name:     name,
				Distance: float32(distance),
				Message:  messageArr,
			})

		}

	}

	answer, err := s.TopSecretService.GetEnemyStarship(&data)
	if err != nil {
		return nil, err
	}

	return answer, nil

}

// Handle for TopSecretSplit godoc
// @Summary get the location and the decoded message
// @Description get the location and the decoded message using the right POST data
// @Tags Meli.QuasarFire
// @Accept json
// @Produce json
// @Param fs_name path string true "First Satellite name"
// @Param ss_name path string true "Seconds Satellite name"
// @Param ts_name path string true "Third Satellite name"
// @Param satellites body aggregate.TopSecretSplitRequest true "Satellites JSON"
// @Success 200 {object} aggregate.TopSecretSplitRequest
// @Router /topsecret_split/{fs_name}/{ss_name}/{ts_name} [post]
func (s *TopSecretSplitHandler) HandlePost(c *gin.Context) (interface{}, error) {

	var (
		satellites = []string{c.Param("fs_name"), c.Param("ss_name"), c.Param("ts_name")}

		requestBody aggregate.TopSecretSplitRequest
		data        aggregate.TopSecretRequest
	)

	// Check: Decode the request HTTP body
	err := c.BindJSON(&requestBody)
	if err != nil {
		return nil, fmt.Errorf("problems when try to deserialize the request body: %s", err.Error())
	}

	// Check: in case has declared no equals number of satellites in the URL and the POST Body
	if len(requestBody.Satellites) != len(satellites) {
		return nil, fmt.Errorf("sorry, we have detected %d elements in the URL and %d in the body, the elements number must be equals", len(satellites), len(requestBody.Satellites))
	}

	for i, name := range satellites {

		data.Satellites = append(data.Satellites, aggregate.TopSecretSatelliteRequest{
			Name:     name,
			Distance: requestBody.Satellites[i].Distance,
			Message:  requestBody.Satellites[i].Message,
		})

	}

	answer, err := s.TopSecretService.GetEnemyStarship(&data)
	if err != nil {
		return nil, err
	}

	return answer, nil

}
