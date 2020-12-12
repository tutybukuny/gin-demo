package models

type Book struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Count       uint   `json:"count"`
	Producer    string `json:"producer"`
}

func (Book) TableName() string {
	return GetTableName()
}

func GetTableName() string {
	return "book"
}
