package apiRoutes

import (
	"encoding/json"
	"fluffy-duck/pkg/utility"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateCem(c *gin.Context, pool *pgxpool.Pool) {
	body, err := io.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to read request body"})
		return
	}
	var jsonData map[string]interface{}
	if err := json.Unmarshal(body, &jsonData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to parse JSON data"})
		return
	}
	timescale := jsonData["timescale"].(string)
	// prevents xxs attack via timescale
	if timescale != "current month" && timescale != "three month rolling" && timescale != "year to date" && timescale != "goal" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid timescale provided",
		})
		return
	}
	osat := jsonData["osat"].(string)
	taste := jsonData["taste"].(string)
	speed := jsonData["speed"].(string)
	ace := jsonData["ace"].(string)
	cleanliness := jsonData["cleanliness"].(string)
	accuracy := jsonData["accuracy"].(string)
	metricSlice := []string{osat, taste, speed, ace, cleanliness, accuracy}
	// prevents xxs attack via other user provided inputs
	for _, value := range metricSlice {
		isNumber := utility.IsStringNumber(value)
		if !isNumber {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "invalid metric provided",
			})
			return
		}
	}
	// push values into db
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
} 