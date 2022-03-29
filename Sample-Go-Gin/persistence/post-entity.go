package persistence

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	ID          uint   `gorm:"primaryKey"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (this *Post) GetID() int {
	return int(this.ID)
}

func (this *Post) SetID(id int) {
	this.ID = uint(id)
}

func (this *Post) GetTitle() string {
	return this.Title
}

func (this *Post) GetDescription() string {
	return this.Description
}
