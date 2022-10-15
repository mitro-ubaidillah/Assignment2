package structs

type Items struct {
	Item_Id     uint   `gorm:"PRIMARY_KEY" json:"lineItemId"`
	Item_Code   string `gorm:"not null" json:"itemCode"`
	Description string `gorm:"type:varchar(191)" json:"description"`
	Quantity    int    `json:"quantity"`
	Order_Id    uint   `json:"order_id"`
}
