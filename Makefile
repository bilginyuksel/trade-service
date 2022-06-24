run:
	go run cmd/main.go

lint:
	golangci-lint run

mockgen:
	mockgen -destination=internal/order/mock/repository.go -package mock github.com/bilginyuksel/trade/internal/order Repository
	mockgen -destination=internal/order/mock/currency_converter.go -package mock github.com/bilginyuksel/trade/internal/order CurrencyConverter