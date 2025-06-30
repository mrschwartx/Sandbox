package catalog

import "errors"

var (
	ErrNotEnoughSeats           = errors.New("no seat available")
	ErrClassSoldOut             = errors.New("class is sold out")
	ErrClassNotAvailableForSale = errors.New("class is not available for sale")
)
