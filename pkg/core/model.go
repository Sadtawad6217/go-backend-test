package core

import "time"

type Exam struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	Title     string    `gorm:"type:text;not null"`
	Content   string    `gorm:"type:text"`
	Published bool      `gorm:"type:boolean;not null;default:false"`
	ViewCount int       `gorm:"type:integer;not null;default:0"`
	CreatedAt time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
	UpdatedAt time.Time `gorm:"type:timestamptz;not null;default:CURRENT_TIMESTAMP"`
}
