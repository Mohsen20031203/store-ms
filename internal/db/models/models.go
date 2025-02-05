package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"` // pending, shipped, delivered, canceled
	Number     uint           `gorm:"foreignKey:OrderID" json:"items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type User struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `json:"name"`
	Email        string         `gorm:"uniqueIndex" json:"email"`
	PasswordHash string         `json:"-"`
	Address      string         `json:"address"`
	PhoneNumber  string         `json:"phone_number"`
	Orders       []Order        ` json:"orders"`
	Cart         Cart           `gorm:"foreignKey:UserID" json:"Cart"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// ------------------------------------------------------
type Product struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Price       float64        `json:"price"`
	Stock       int            `json:"stock"`
	CategoryID  []uint         `json:"category_id"`
	BrandID     []uint         `json:"brand_id"`
	Reviews     []Review       `gorm:"foreignKey:ProductID" json:"reviews"`
	ImageURLs   string         `json:"image_urls"`
	Rating      float64        `json:"rating"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
}

type Category struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	ParentID *uint     `json:"parent_id"`
	Products []Product `gorm:"foreignKey:CategoryID" json:"products"`
}

type Brand struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Name     string    `json:"name"`
	Products []Product `gorm:"foreignKey:BrandID" json:"products"`
}
type Payment struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Order         *Order    `gorm:"foreignKey:OrderID"`
	User          User      `gorm:"foreignKey:UserID"`
	Amount        float64   `json:"amount"`
	Status        string    `json:"status"` // pending, successful, failed
	TransactionID string    `gorm:"uniqueIndex" json:"transaction_id"`
	CreatedAt     time.Time `json:"created_at"`
}

type Cart struct {
	ID        uint       `gorm:"primaryKey" json:"id"`
	UserID    uint       `json:"user_id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ID       uint    `gorm:"primaryKey" json:"id"`
	CartID   uint    `json:"cart_id"`
	Product  Product `gorm:"foreignKey:ProductID"`
	Quantity int     `json:"quantity"`
	Price    float64 `json:"price"`
}

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	User      User      `json:"user"`
	Rating    float64   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}
