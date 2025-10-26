package service

import (
	"car_project/internal/entities"
	"errors"
	"time"

	"gorm.io/gorm"
)

type WalletService struct{}

func (w *WalletService) DebitWallet(tx *gorm.DB, userID, orderID int, amount float64) error {
	var wallet entities.Wallet
	if err := tx.Where("user_id = ?", userID).First(&wallet).Error; err != nil {
		return errors.New("wallet not found")
	}

	if wallet.Balance < amount {
		return errors.New("insufficient balance")
	}

	wallet.Balance -= amount
	now := time.Now()
	wallet.UpdatedAt = now.Format("2006-01-02 15:04:05")
	if err := tx.Save(&wallet).Error; err != nil {
		return err
	}

	// créer transaction wallet
	walletTx := &entities.WalletTransaction{
		WalletID:        wallet.ID,
		TransactionType: "debit",
		Amount:          amount,
		Reference:       "order_" + string(orderID),
		Description:     "Paiement commande",
		CreatedAt:       time.Now(),
	}

	return tx.Create(walletTx).Error
}
