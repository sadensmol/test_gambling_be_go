package tests_test

import (
	"math"
	"testing"

	"context"
	"net/http"
	"net/http/cookiejar"
	"time"

	v1 "github.com/sadensmol/test_gambling_be_go/api/proto/v1"
	"github.com/stretchr/testify/suite"

	"github.com/steinfletcher/apitest"
	jsonpath "github.com/steinfletcher/apitest-jsonpath"
	"google.golang.org/grpc"
	"google.golang.org/grpc/benchmark"
	"google.golang.org/grpc/credentials/insecure"
)

type WalletTestSuite struct {
	suite.Suite
}

const (
	testUserID = 1
)

func TestWalletTestSuite(t *testing.T) {
	suite.Run(t, new(WalletTestSuite))

}

func (s *WalletTestSuite) TestDepositThenWithdrawSuccess() {
	clientConn := benchmark.NewClientConn("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := v1.NewWalletServiceClient(clientConn)
	depositAmount := 111
	withdrawAmount := 12

	r, err := client.GetByUserId(context.Background(), &v1.GetWalletBalanceRequest{UserID: testUserID})
	s.Require().NoError(err)
	s.Require().NotNil(r)

	startAmount := r.Balance

	r2, err := client.Deposit(context.Background(), &v1.DepositRequest{UserID: testUserID, Amount: int64(depositAmount)})
	s.Require().NoError(err)
	s.Require().NotNil(r2)

	r3, err := client.Withdraw(context.Background(), &v1.WithdrawRequest{UserID: testUserID, Amount: int64(withdrawAmount)})
	s.Require().NoError(err)
	s.Require().NotNil(r3)

	r, err = client.GetByUserId(context.Background(), &v1.GetWalletBalanceRequest{UserID: testUserID})
	s.Require().NoError(err)
	s.Require().NotNil(r)
	s.Require().Equal(startAmount+int64(depositAmount)-int64(withdrawAmount), r.Balance)

}
func (s *WalletTestSuite) TestDepositAndWithdrawalSuccessHttp() {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	apitest.New().
		EnableNetworking(cli).
		Post("http://localhost:8090/api/wallet/deposit").
		Bodyf(`{"user_id":%d,"amount":%d}`, testUserID, 123).
		Expect(s.T()).
		Body("{}").
		Status(http.StatusOK).
		End()

	apitest.New().
		EnableNetworking(cli).
		Post("http://localhost:8090/api/wallet/withdraw").
		Bodyf(`{"user_id":%d,"amount":%d}`, testUserID, 12).
		Expect(s.T()).
		Body("{}").
		Status(http.StatusOK).
		End()
}

func (s *WalletTestSuite) TestWithdrawNotEnoughMoneyError() {
	clientConn := benchmark.NewClientConn("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := v1.NewWalletServiceClient(clientConn)

	_, err := client.Withdraw(context.Background(), &v1.WithdrawRequest{UserID: testUserID, Amount: math.MaxInt64})
	s.Require().Error(err, "not enough money")
}

func (s *WalletTestSuite) TestWithdrawNotEnoughMoneyErrorHttp() {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	apitest.New().
		EnableNetworking(cli).
		Post("http://localhost:8090/api/wallet/withdraw").
		Bodyf(`{"user_id":%d,"amount":%d}`, testUserID, math.MaxInt64).
		Expect(s.T()).
		Assert(jsonpath.Contains("message", "not enough money")).
		Status(http.StatusInternalServerError).
		End()

}

func (s *WalletTestSuite) TestGetBalanceHttp() {
	cookieJar, _ := cookiejar.New(nil)
	cli := &http.Client{
		Timeout: time.Second * 1,
		Jar:     cookieJar,
	}

	apitest.New().
		EnableNetworking(cli).
		Getf("http://localhost:8090/api/wallet/balance/%d", testUserID).
		Expect(s.T()).
		Assert(jsonpath.Present("balance")).
		Status(http.StatusOK).
		End()

}
