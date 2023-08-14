package controllers

import (
	"context"
	"fmt"

	v1 "github.com/sadensmol/test_gambling_be_go/api/proto/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type WalletHandler struct {
	walletService IWalletService
}

type IWalletService interface {
	Deposit(userID int64, amount int64) error
	Withdraw(userID int64, amount int64) error
	GetBalance(userID int64) (int64, error)
}

func NewWalletHandler(walletService IWalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func (h *WalletHandler) Deposit(ctx context.Context, req *v1.DepositRequest) (*emptypb.Empty, error) {
	err := h.walletService.Deposit(req.UserID, req.Amount)

	if err != nil {
		_ = fmt.Errorf("Error occurred %w", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (h *WalletHandler) Withdraw(ctx context.Context, req *v1.WithdrawRequest) (*emptypb.Empty, error) {
	err := h.walletService.Withdraw(req.UserID, req.Amount)

	if err != nil {
		_ = fmt.Errorf("Error occurred %w", err)
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
func (h *WalletHandler) GetByUserId(ctx context.Context, req *v1.GetWalletBalanceRequest) (*v1.GetWalletBalanceResponse, error) {
	balance, err := h.walletService.GetBalance(req.UserID)

	if err != nil {
		_ = fmt.Errorf("Error occurred %w", err)
		return nil, err
	}

	return &v1.GetWalletBalanceResponse{Balance: balance}, nil
}
