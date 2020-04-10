package models

import (
	"time"

	"github.com/MarErm27/uadmin"
)

// Bid model ...
type BiddingHistory struct {
	uadmin.Model
	Name      string  `uadmin:"required;search;multilingual"`
	Auction   Auction // <-- Category Model
	AuctionID uint    // <-- CategoryID

	Participant   Participant
	ParticipantID uint

	Bid     string
	BidTime time.Time
	BestBid bool
}
