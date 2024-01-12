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
	ID         uint   `gorm:"primaryKey"`
	Name       string `gorm:"not null"`
	Price      float64
	Units      int32
	CategoryID uint     // Llave foránea a Category
	Category   Category `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
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
	ID              uint          `gorm:"primaryKey"`
	UserID          uint          // Llave foránea a User
	User            User          `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	PaymentMethodID uint          // Llave foránea a PaymentMethod
	PaymentMethod   PaymentMethod `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TotalAmount     float64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

type OrderProduct struct {
	ID        uint    `gorm:"primaryKey"`
	OrderID   uint    // Llave foránea a Order
	Order     Order   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	ProductID uint    // Llave foránea a Product
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
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
