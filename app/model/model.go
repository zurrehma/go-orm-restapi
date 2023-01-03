package model

import (
	"time"

	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title string `gorm:"unique;not null" json:"title"`
	Tasks []Task `gorm:"foreignKey:ProjectID" json:"tasks"`
}

type Task struct {
	gorm.Model
	Title     string    `gorm:"unique" json:"title"`
	Priority  string    `gorm:"type:ENUM('0', '1', '2', '3');default:'0'" json:"priority"`
	Deadline  time.Time `gorm:"default:null" json:"deadline"`
	Done      bool      `json:"done"`
	ProjectID uint      `json:"project_id"`
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Project{}, &Task{})
	return db
}
