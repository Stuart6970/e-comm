package models

import (
	"errors"
	"html"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/shopspring/decimal"
)

type CatalogItem struct {
	ID              uint64          `gorm:"primary_key;auto_increment" json:"id"`
	Name            string          `gorm:"size:255;not null;" json:"name"`
	Description     string          `gorm:"size:500;" json:"description"`
	Price           decimal.Decimal `gorm:"type:decimal(10,2);" json:"price"`
	PictureFileName string          `gorm:"size:255;not null;" json:"picture_file_name"`
	PictureUri      string          `gorm:"size:500;not null;" json:"picture_uri"`
	AvailableStock  uint32          `gorm:"not null" json:"available_stock"`
	CreatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time       `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *CatalogItem) Prepare() {
	c.ID = 0
	c.Name = html.EscapeString(strings.TrimSpace((c.Name)))
	c.Description = html.EscapeString(strings.TrimSpace((c.Description)))
	c.PictureFileName = html.EscapeString(strings.TrimSpace((c.PictureFileName)))
	c.PictureUri = html.EscapeString(strings.TrimSpace((c.PictureUri)))
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()
}

func (c *CatalogItem) Validate() error {
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Description == "" {
		return errors.New("description is required")
	}
	if c.PictureFileName == "" {
		return errors.New("PictureFileName is required")
	}
	if c.PictureUri == "" {
		return errors.New("PictureUri is required")
	}
	return nil
}

func (c *CatalogItem) SaveCatalogItem(db *gorm.DB) (*CatalogItem, error) {
	var err error
	err = db.Debug().Model(&CatalogItem{}).Create(&c).Error
	if err != nil {
		return &CatalogItem{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&CatalogItem{}).Where("ID = ?", &c.ID).Take(&c).Error
		if err != nil {
			return &CatalogItem{}, err
		}
	}
	return c, nil
}

func (c *CatalogItem) FindAllCatalogItems(db *gorm.DB) (*[]CatalogItem, error) {
	var err error
	catalogItems := []CatalogItem{}
	err = db.Debug().Model(&CatalogItem{}).Limit(100).Find(&catalogItems).Error
	if err != nil {
		return &[]CatalogItem{}, err
	}
	return &catalogItems, err
}

func (c *CatalogItem) FindCatalogItemById(db *gorm.DB, cid uint64) (*CatalogItem, error) {
	var err error
	err = db.Debug().Model(&CatalogItem{}).Where("id = ?", cid).Take(&c).Error
	if err != nil {
		return &CatalogItem{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&CatalogItem{}).Where("id = ?", cid).Take(&c).Error
		if err != nil {
			return &CatalogItem{}, err
		}
	}
	return c, nil
}

func (c *CatalogItem) UpdateCatalogItem(db *gorm.DB) (*CatalogItem, error) {
	var err error
	err = db.Debug().Model(&CatalogItem{}).Updates(CatalogItem{Name: c.Name, Description: c.Description,
		Price: c.Price, AvailableStock: c.AvailableStock, PictureFileName: c.PictureFileName,
		PictureUri: c.PictureUri, UpdatedAt: time.Now()}).Error
	if err != nil {
		return &CatalogItem{}, err
	}
	if c.ID != 0 {
		err = db.Debug().Model(&CatalogItem{}).Where("id = ?", c.ID).Take(&c).Error
		if err != nil {
			return &CatalogItem{}, err
		}
	}
	return c, nil
}

func (c *CatalogItem) DeleteCatalogItem(db *gorm.DB, cid uint64) (int64, error) {

	db = db.Debug().Model(&CatalogItem{}).Where("id = ?", cid).Take(&CatalogItem{}).Delete(&CatalogItem{})
	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return 0, errors.New("post not found")
		}
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
