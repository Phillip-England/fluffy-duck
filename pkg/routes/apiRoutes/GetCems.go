package apiRoutes

import (
	"fluffy-duck/pkg/database"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func GetCems(c *gin.Context, pool *pgxpool.Pool) {
	scores, err := database.GetCemScores(pool)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "internal server error",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"cems": scores,
	})
} 