package models

type Cart struct {
	Id        int64 `json:"id" gorm:"primaryKey"`
	Quantity  int32 `json:"quantity"`
	ProductId int64 `json:"product_id"`
	UserId    int64 `json:"user_id"`
}
