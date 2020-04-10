package models

import (
	"github.com/MarErm27/uadmin"
)

// AuctionTable model ...
type AuctionTable struct {
	uadmin.Model
	Name          string
	Auction       Auction
	AuctionID     uint
	Participant   Participant
	ParticipantID uint
	FallLimit     string
	PlanFallLimit string
	Fall          string
}
