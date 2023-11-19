package pageRoutes

import (
	"fluffy-duck/pkg/database"
	"fluffy-duck/pkg/middleware"
	"fluffy-duck/pkg/render"
	"fluffy-duck/pkg/utility"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
)

func UpdateCemPage(c *gin.Context, pool *pgxpool.Pool) {
	userType, _ := middleware.Auth(c)
	if userType == "team" {
		c.Redirect(301, "/team")
		return
	}
	timescale := c.Query("timescale")
	if timescale == "" {
		c.Redirect(301, "/admin/cem/update?timescale=current month")
		return
	}
	allowedTimescales := []string{"current month", "three month rolling", "year to date", "annual goal"}
	timescaleAllowed := utility.IsStringInSlice(allowedTimescales, timescale)
	if !timescaleAllowed {
		c.Redirect(301, "/admin/cem/update")
		return
	}
	scores, err := database.GetCemScoresByTimescale(pool, timescale)
	if err != nil {
		fmt.Println(err)
	}
	c.Data(200, "text/html", []byte(render.UpdateCemPage(c, scores)))
}