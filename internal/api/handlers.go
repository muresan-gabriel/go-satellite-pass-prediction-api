package api

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type PassHandler struct{}

func NewPassHandler() *PassHandler {
	return &PassHandler{}
}

type getPassesRequest struct {
	SatelliteID  string  `json:"satelliteId" binding:"required"`
	Latitude     float64 `json:"latitude" binding:"required"`
	Longitude    float64 `json:"longitude" binding:"required"`
	Elevation    float64 `json:"elevation"`
	Start        string  `json:"start" binding:"required"`
	End          string  `json:"end" binding:"required"`
	MinElevation float64 `json:"minElevation"`
}

func extractTime(req getPassesRequest) (*time.Time, *time.Time, error) {
	start, err := time.Parse(time.RFC3339, req.Start)
	end, err := time.Parse(time.RFC3339, req.End)

	if err != nil {
		return nil, nil, err
	}

	return &start, &end, nil
}

func validateDates(start, end time.Time) error {
	if end.Before(start) {
		return fmt.Errorf("end time must be after start time")
	}
	return nil
}

func (h *PassHandler) GetPasses(c *gin.Context) {
	var req getPassesRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	start, end, err := extractTime(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid time format",
		})
		return
	}

	if err := validateDates(*start, *end); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, []gin.H{
		{
			"aos":              start.Add(10 * time.Minute),
			"los":              start.Add(18 * time.Minute),
			"max_elevation":    42.5,
			"duration_seconds": 480,
			"visibility":       "NIGHT",
		},
	})
}
