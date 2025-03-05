package models

import (
	"time"

	"gorm.io/gorm"
)

type Order struct {
	ID         uint           `gorm:"primaryKey" json:"id"`
	TotalPrice float64        `json:"total_price"`
	Status     string         `json:"status"` // pending, shipped, delivered, canceled
	Number     uint           `gorm:"primaryKey" json:"items"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
}

type Review struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UserId    uint      `json:"user_id"`
	Rating    float64   `json:"rating"`
	Comment   string    `json:"comment"`
	CreatedAt time.Time `json:"created_at"`
}

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

type User struct {
	gorm.Model   `json:"-"`
	ID           uint           `gorm:"primaryKey" json:"id"`
	Name         string         `gorm:"type:varchar(500)" json:"name"`
	Email        string         `gorm:"type:varchar(500)" json:"email"`
	PasswordHash string         `gorm:"type:varchar(500)" json:"-"`
	Address      string         `gorm:"type:varchar(500)" json:"address"`
	PhoneNumber  string         `gorm:"type:varchar(500)" json:"phone_number"`
	Orders       []Order        `gorm:"foreignKey:UserID" json:"orders"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
}

// ------------------------------------------------------

type Category struct {
	ID       uint   `gorm:"primaryKey" json:"id"`
	Name     string `json:"name"`
	ParentID *uint  `json:"parent_id"`
}

type Brand struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
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
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	Items     []CartItem `gorm:"foreignKey:CartID" json:"items"`
}

type CartItem struct {
	ProductId uint    `gorm:"foreignKey:ProductID"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
}
