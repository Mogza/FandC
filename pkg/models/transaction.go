package models

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	TokenID         uint `gorm:"index:idx_transaction_token_id"`
	BuyerID         uint `gorm:"index:idx_transaction_buyer_id"`
	SellerID        uint `gorm:"index:idx_transaction_seller_id"`
	Amount          float64
	TransactionDate string
}
