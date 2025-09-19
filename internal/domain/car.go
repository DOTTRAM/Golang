package domain

type Car struct {
	Id    int    `gorm:"primaryKey;" json:"id"`
	Name  string `gorm:"type:varchar(50)" json:"name"`
	Year  int    `json:"year"`
	Price int    `json:"price"`
}
