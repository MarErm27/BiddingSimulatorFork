package models

import (
	"github.com/MarErm27/uadmin"
)

// Auction model ...
type RealBidders struct {
	uadmin.Model
	Name            string
	FallLimit       float64
	LastOffer       float64
	CurrentPosition int
}
