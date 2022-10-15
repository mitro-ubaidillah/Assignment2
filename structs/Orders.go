package structs

import (
	"time"
)

type Orders struct {
	Order_Id      uint      `gorm:"PRIMARY_KEY;"`
	Customer_Name string    `gorm:"type:varchar(191)" json:"customerName"`
	Ordered_At    time.Time `json:"orderedAt"`
	Item          []Items   `gorm:"foreignKey:order_id" json:"items"`
}
