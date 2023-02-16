package models

type Verdict struct {
	ID          string `json:"verdict_id" gorm:"primary_key"`
	Description string `json:"description" gorm:"not null;unique"`
}
