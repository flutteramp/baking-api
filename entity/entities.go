package entity

import (
	"time"
)

type Recipe struct {
	ID          uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID      uint   `json:"userid"`
	Title       string `json:"title" gorm:"type:varchar(255);not null"`
	Duration    string `json:"duration" gorm:"type:varchar(255);not null"`
	Servings    int    `json:"servings"`
	RecipeUser  uint   `json:"recipeuserid" gorm:"type:varchar(255);not null`
	ImageUrl    string `json:"imageUrl" gorm:"type:varchar(255);not null"`
	Comments    []Comment
	Ingredients []Ingredient
	Steps       []Step
}

type User struct {
	ID       uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	FullName string `json:"fullname" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"password" gorm:"type:varchar(255);not null"`
	Recipes  []Recipe
}

type Comment struct {
	ID        uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint   `json:"userid"`
	UserName  string `json:"username" gorm:"type:varchar(255);not null`
	RecipeID  uint   `json:"recipeid" gorm:"type:varchar(255);not null`
	Message   string `json:"message" gorm:"type:varchar(255);not null"`
	CreatedAt time.Time
}

type Ingredient struct {
	ID         uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RecipeID   uint   `json:"recipeid" `
	Title      string `json:"title" gorm:"type:varchar(255);not null`
	Quantity   uint   `json:"quantity" gorm:"DEFAULT:0"`
	Measurment string `json:"measurment" gorm:"type:varchar(255);not null`
}
type Step struct {
	ID        uint   `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RecipeID  uint   `json:"recipeid" `
	Direction string `json:"direction" gorm:"type:varchar(255);not null`
}
