package order

import "time"

type Order struct {
	ID     string
	Bidder string
	Bid    string
	Ask    string
	Amount float64

	At time.Time
}
