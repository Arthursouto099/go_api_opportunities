package schemas

import (
	"gorm.io/gorm"
)

type Opening struct {
	// propriedades do gorm.Model
	// exc: CreatedAt UpdatedAt Id
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}
