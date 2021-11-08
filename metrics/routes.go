package metrics

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jprice8/website-analytics/shared"
)

func PageViewsRegister(router *gin.RouterGroup) {
	router.GET("/", hitsList)
	// router.POST("/", pageViewCreate)
}

func hitsList(c *gin.Context) {
	hits, err := getHits()
	if err != nil {
		c.JSON(http.StatusNotFound, shared.NewError("hitsList", errors.New("Invalid param")))
		return
	}
	serializer := HitsSerializer{c, hits}
	c.JSON(http.StatusOK, gin.H{"hits": serializer.Response()})
}

// func createHit(c *gin.Context) {
// 	if err := SaveOne(&pageViewModelValidator.pageViewModel); err != nil {
// 		c.JSON(http.StatusUnprocessableEntity, shared.NewError("database", err))
// 		return
// 	}
// 	serializer := PageViewSerializer{c, pageViewModelValidator.pageViewModel}
// 	c.JSON(http.StatusCreated, gin.H{"pageView": serializer.Response()})
// }
