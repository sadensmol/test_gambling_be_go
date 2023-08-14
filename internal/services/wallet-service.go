package services

import (
	"sync"

	"github.com/sadensmol/test_gambling_be_go/internal/domain"
)

type WalletService struct {
	wallets map[int64]int64
	wsMx    sync.Mutex
}

func NewWalletService() *WalletService {
	return &WalletService{wallets: make(map[int64]int64)}
}

func (s *WalletService) Deposit(userID int64, amount int64) error {
	s.wsMx.Lock()
	defer s.wsMx.Unlock()
	s.wallets[userID] += amount
	return nil
}

func (s *WalletService) Withdraw(userID int64, amount int64) error {
	s.wsMx.Lock()
	defer s.wsMx.Unlock()

	wa := s.wallets[userID]
	if wa < amount {
		return domain.ErrNotEnoughMoney
	}

	s.wallets[userID] -= amount
	return nil
}

func (s *WalletService) GetBalance(userID int64) (int64, error) {
	s.wsMx.Lock()
	defer s.wsMx.Unlock()

	wa := s.wallets[userID]

	return wa, nil
}
