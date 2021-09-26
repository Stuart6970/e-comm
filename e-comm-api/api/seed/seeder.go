package seed

import (
	"log"
	"time"

	"github.com/Stuart6970/e-comm-api/api/models"
	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

var catalogItems = []models.CatalogItem{
	{
		Name:            "BLONDE TOP",
		Description:     "Top med blondedetalje fra JDY",
		Price:           decimal.NewFromFloat(125.95),
		PictureFileName: "15236091_EcruOlive_001_ProductLarge.jpg",
		PictureUri:      "https://www.only.com/dw/image/v2/ABBT_PRD/on/demandware.static/-/Sites-pim-catalog/default/dwe5aefa57/pim-static/large/",
		AvailableStock:  27,
		CreatedAt:       time.Now(),
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.CatalogItem{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.CatalogItem{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i := range catalogItems {
		err = db.Debug().Model(&models.CatalogItem{}).Create(&catalogItems[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}
}
