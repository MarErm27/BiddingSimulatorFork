package models

import (
	"time"

	"github.com/MarErm27/uadmin"
)

// Offer model ...
type Offer struct {
	uadmin.Model
	Name           string
	AuctionTable   AuctionTable // <-- Category Model
	AuctionTableID uint         // <-- CategoryID
	Bid            string
	BidTime        time.Time
	BestBid        bool
}
