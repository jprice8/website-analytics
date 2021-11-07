package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jprice8/website-analytics/shared"
	"github.com/jprice8/website-analytics/metrics"
	"gorm.io/gorm"
)

// album represents data about a record album.
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&metrics.PageViewModel{})
	db.AutoMigrate(&Album{})
}

func main() {

	db := shared.Init()
	Migrate(db)

	db.Create(&Album{ID: "4", Title: "Easy Wind", Artist: "Grateful Dead", Price: 1.99})

	router := gin.Default()

	v1 := router.Group("/api")
	metrics.PageViewsRegister(v1.Group("/metrics"))


	router.Run("localhost:8080")
}
