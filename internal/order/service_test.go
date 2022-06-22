package order_test

import (
	"context"
	"testing"
	"time"

	"github.com/bilginyuksel/trade/internal/order"
	"github.com/bilginyuksel/trade/internal/order/mock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestBid(t *testing.T) {
	mockRepository := mock.NewMockRepository(gomock.NewController(t))
	mockCurrencyConverter := mock.NewMockCurrencyConverter(gomock.NewController(t))

	sampleOrders := []order.Order{{
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.23,
	}, {
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.33,
	}, {
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.85,
	}, {
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 1.12,
	}}

	mockRepository.EXPECT().
		Filter(gomock.Any(), "BTC", "ETH").
		Return(sampleOrders, nil)
	mockCurrencyConverter.EXPECT().
		Convert(gomock.Any(), "ETH", "BTC", 5.84).
		Return(0.623813, nil)
		// 0,063813
		// 0.85 - 0.063813 = 0.786187

	testOrder := &order.Order{
		Bid:    "ETH",
		Ask:    "BTC",
		Amount: 5.84,
		At:     time.Now().UTC(),
	}

	expectedOrders := []order.Order{{
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.23,
	}, {
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.33,
	}, {
		Bid:    "BTC",
		Ask:    "ETH",
		Amount: 0.063813,
	}}

	service := order.NewService(mockRepository, mockCurrencyConverter)
	orders, err := service.Bid(context.TODO(), testOrder)

	assert.Nil(t, err)
	assert.Equal(t, expectedOrders, orders)
}
