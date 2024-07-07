// models/page_visit.go
package models

import (
	"time"

	"gorm.io/gorm"
)

type PageVisit struct {
	ID        uint      `gorm:"primaryKey"`
	Page      string    `json:"page"`
	VisitTime time.Time `json:"visitTime"`
	IP        string    `json:"ip"`
}

func (p *PageVisit) Save(db *gorm.DB) error {
	return db.Create(p).Error
}

func GetVisitStats(db *gorm.DB) ([]struct {
	Page  string `json:"page"`
	Count int    `json:"count"`
}, error) {
	var stats []struct {
		Page  string `json:"page"`
		Count int    `json:"count"`
	}
	err := db.Model(&PageVisit{}).Select("page, count(*) as count").Group("page").Find(&stats).Error
	return stats, err
}
