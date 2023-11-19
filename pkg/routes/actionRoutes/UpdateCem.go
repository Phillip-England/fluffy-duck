package actionRoutes

import (
	"fluffy-duck/pkg/database"
	"fluffy-duck/pkg/models"
	"fluffy-duck/pkg/utility"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateCem(c *gin.Context, pool *pgxpool.Pool) {
	timescale := c.Query("timescale")
	osat := c.PostForm("osat")
	taste := c.PostForm("taste")
	speed := c.PostForm("speed")
	ace := c.PostForm("ace")
	cleanliness := c.PostForm("cleanliness")
	accuracy := c.PostForm("accuracy")
	cemScore := models.CemScore {
		Timescale: timescale,
		Osat: utility.StringToInt(osat),
		Taste: utility.StringToInt(taste),
		Speed: utility.StringToInt(speed),
		Ace: utility.StringToInt(ace),
		Cleanliness: utility.StringToInt(cleanliness),
		Accuracy: utility.StringToInt(accuracy),
	}
	database.UpdateCemScoreByTimescale(pool, &cemScore)
	c.Redirect(301, "/admin/cem/update?timescale=" + cemScore.Timescale)
}