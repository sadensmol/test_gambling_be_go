package services

import (
	"testing"

	"github.com/sadensmol/test_gambling_be_go/internal/domain"
)

func TestWalletService(t *testing.T) {
	testCases := []struct {
		name             string
		depositAmount    int64
		withdrawAmount   int64
		expectedError    error
		expectedBalance  int64
		expectedNotFound bool
	}{
		{
			name:            "Valid Deposit and Withdraw",
			depositAmount:   100,
			withdrawAmount:  50,
			expectedError:   nil,
			expectedBalance: 50,
		},
		{
			name:            "Not Enough Money to Withdraw",
			depositAmount:   100,
			withdrawAmount:  150,
			expectedError:   domain.ErrNotEnoughMoney,
			expectedBalance: 100,
		},
	}

	userID := int64(1)

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			// Initialize WalletService for testing
			service := NewWalletService()

			// Perform actions if needed
			if tc.depositAmount > 0 {
				err := service.Deposit(userID, tc.depositAmount)
				if err != nil {
					t.Errorf("Deposit failed: %v", err)
				}
			}
			if tc.withdrawAmount > 0 {
				err := service.Withdraw(userID, tc.withdrawAmount)
				if err != tc.expectedError {
					t.Errorf("Expected error: %v, Got error: %v", tc.expectedError, err)
				}
			}

			// Balance check
			balance, err := service.GetBalance(userID)

			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if balance != tc.expectedBalance {
				t.Errorf("Expected balance: %d, Got balance: %d", tc.expectedBalance, balance)
			}
		})
	}
}
