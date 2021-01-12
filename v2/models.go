import()

// task structure
type Task struct {
	gorm.Model
	id  uint 	`gorm:"primaryKey"`
	Name string `json:"name"`
	Done bool    `json:"done"`
}
