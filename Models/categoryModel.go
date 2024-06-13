package models

type Category struct {
	CategoryID int    `json:"category_id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT; primaryKey"`
	Name       string `json:"name"`
	CreatedAt  string `json:"created_at"`
	UpdatedAt  string `json:"updated_at"`
}
