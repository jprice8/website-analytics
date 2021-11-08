package metrics

import (
	"fmt"

	"github.com/jprice8/website-analytics/shared"
	"gorm.io/gorm"
)

type Hit struct {
	gorm.Model
	Url		string		`json:"url"` 
}

func getHits() ([]Hit, error) {
	db := shared.GetDB()
	var hits []Hit
	fmt.Println(&hits)
	err := db.Find(&hits).Error
	return hits, err
}

func SaveOne(data interface{}) error {
	db := shared.GetDB()
	err := db.Save(data).Error
	return err
}