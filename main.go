package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jprice8/website-analytics/metrics"
	"github.com/jprice8/website-analytics/shared"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func listHits(c *gin.Context) {
	// Query for data
	var hits []metrics.Hit
	err := db.Find(&hits).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewError("listHits", errors.New("problem with everything")))
		return
	}

	jsonEncoded, err := json.Marshal(&hits)
	if err != nil {
		c.JSON(http.StatusBadRequest, shared.NewError("jsonEncoded", errors.New("problem with encoding the json")))
	}

	c.JSON(http.StatusOK, gin.H{"hits": jsonEncoded})
}

func main() {
	db, err = gorm.Open(sqlite.Open("test1.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database :(")
	}

	// Migrate the schema
	db.AutoMigrate(&metrics.Hit{})

	// Create sample url visit
	db.Create(&metrics.Hit{Url: "jprice.io"})
	// Find all page views.
	var hits []metrics.Hit
	db.Find(&hits)
	fmt.Println(hits)

	// Output the users from the DB json encoded
	jsonEncoded, _ := json.Marshal(&hits)
	fmt.Println(jsonEncoded)

	router := gin.Default()

	v1 := router.Group("/api")
	v1.GET("/", listHits)
	// metrics.PageViewsRegister(v1.Group("/metrics"))

	router.Run("localhost:8080")
}
