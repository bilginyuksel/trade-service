package order

import (
	"context"
	"errors"
	"log"
)

// TODO: return this error when there is no order in the order book
var ErrNoOrders = errors.New("no orders")

type (
	Repository interface {
		Insert(ctx context.Context, order *Order) error
		Get(ctx context.Context, id string) (*Order, error)
		// Filter returns all orders that match the given bid and ask
		Filter(ctx context.Context, bid, ask string) ([]Order, error)
	}

	CurrencyConverter interface {
		Convert(ctx context.Context, from, to string, amount float64) (float64, error)
	}
)

type Service struct {
	repo Repository
	cc   CurrencyConverter
}

func NewService(repo Repository, cc CurrencyConverter) *Service {
	return &Service{
		repo: repo,
		cc:   cc,
	}
}

// Filter returns all orders that match the given bid and ask
// returns orders in descending order of at time
func (s *Service) Filter(ctx context.Context, bid, ask string) ([]Order, error) {
	return s.repo.Filter(ctx, bid, ask)
}

// Bid an order to the order book
// Order book will execute the order async if there is a match
// If there is no match, the order will be placed in the order book
func (s *Service) Bid(ctx context.Context, order *Order) ([]Order, error) {
	potentialOrders, err := s.Filter(ctx, order.Ask, order.Bid)
	if err != nil {
		log.Printf("error getting orders: %v\n", err)
		return nil, err
	}

	orderAskAmount, err := s.cc.Convert(ctx, order.Bid, order.Ask, order.Amount)
	if err != nil {
		log.Printf("error converting order amount: %v\n", err)
		return nil, err
	}

	return s.satisfyOrder(ctx, orderAskAmount, potentialOrders), nil
}

func (s *Service) satisfyOrder(ctx context.Context, orderAskAmount float64, potentialOrders []Order) (orders []Order) {
	for i := 0; i < len(potentialOrders) && orderAskAmount >= potentialOrders[i].Amount; i++ {
		orderAskAmount -= potentialOrders[i].Amount
		orders = append(orders, potentialOrders[i])
	}

	return
}
