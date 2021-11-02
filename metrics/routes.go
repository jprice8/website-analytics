package metrics

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jprice8/website-analytics/shared"
)

func PageViewsRegister(router *gin.RouterGroup) {
	router.GET("/", pageViewsList)
	router.POST("/", pageViewCreate)
}

func pageViewsList(c *gin.Context) {
	pageViews, err := getAllPageViews()
	if err != nil {
		c.JSON(http.StatusNotFound, shared.NewError("pageViews", errors.New("Invalid param")))
		return
	}
	serializer := PageViewsSerializer{c, pageViews}
	c.JSON(http.StatusOK, gin.H{"pageViews": serializer.Response()})
}

func pageViewCreate(c *gin.Context) {
	pageViewModelValidator := NewPageViewModelValidator()
	if err := SaveOne(&pageViewModelValidator.pageViewModel); err != nil {
		c.JSON(http.StatusUnprocessableEntity, shared.NewError("database", err))
		return
	}
	serializer := PageViewSerializer{c, pageViewModelValidator.pageViewModel}
	c.JSON(http.StatusCreated, gin.H{"pageView": serializer.Response()})
}
