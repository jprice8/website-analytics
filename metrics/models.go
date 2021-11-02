package metrics

import (
	"github.com/jprice8/website-analytics/shared"
	"gorm.io/gorm"
)

type PageViewModel struct {
	gorm.Model
	url			string 
}

func getAllPageViews() ([]PageViewModel, error) {
	db := shared.GetDB()
	var pageViews []PageViewModel
	err := db.Find(&pageViews).Error
	return pageViews, err
}

func SaveOne(data interface{}) error {
	db := shared.GetDB()
	err := db.Save(data).Error
	return err
}