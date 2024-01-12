package models

import "time"

type User struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Email     string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Product struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	Price     float64
	Units     int32
	Category  Category
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Currency struct {
	ID              uint   `gorm:"primaryKey"`
	Name            string `gorm:"not null"`
	Symbol          string `gorm:"not null"`
	DollarTranslate float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type Order struct {
	ID            uint `gorm:"primaryKey"`
	User          User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID        uint
	PaymentMethod PaymentMethod
	TotalAmount   float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type OrderProduct struct {
	ID        uint  `gorm:"primaryKey"`
	Order     Order `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	OrderID   uint
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID uint
	Quantity  int
}

type Category struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type PaymentMethod struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
