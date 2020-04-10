package models

import (
	"time"

	"github.com/MarErm27/BiddingSimulatorFork/MarErm"
	"github.com/MarErm27/uadmin"
)

// Auction model ...
type Auction struct {
	uadmin.Model
	Name                            string `uadmin:"required;search;multilingual"`
	Buyer                           string `uadmin:"required;search;multilingual"`
	PublicationDate                 time.Time
	ApplicationDeadline             time.Time
	BiddingStart                    time.Time
	BiddingFinish                   time.Time
	Status                          string `uadmin:"search;multilingual"`
	InitialPrice                    string `uadmin:"search;pattern:^[1-9][0-9]*([.][0-9]{2}|)*$;pattern_msg:Your input must be a number."`
	HalfPercentOfInitialPrice       string `uadmin:"pattern:^[1-9][0-9]*([.][0-9]{2}|)*$;pattern_msg:Your input must be a number."`
	FivePercentsOfInitialPrice      string `uadmin:"pattern:^[1-9][0-9]*([.][0-9]{2}|)*$;pattern_msg:Your input must be a number."`
	CurrentPrice                    string `uadmin:"pattern:^[1-9][0-9]*([.][0-9]{2}|)*$;pattern_msg:Your input must be a number."`
	CurrentFall                     string `uadmin:"pattern:^[1-9][0-9]*([.][0-9]{2}|)*$;pattern_msg:Your input must be a number."`
	TimeToFinish                    time.Time
	CurrentParticipantWithBestOffer string
	CurrentBestOfferTime            time.Time
	TotalOffersSubmitted            int
	EstimatedDatesAndEndTime        time.Time
}

func (a *Auction) Save() {
	// Multiply the Number and the Cost to get the value of the Sum
	a.CurrentPrice = a.InitialPrice
	initialPrice := MarErm.PriceToInt64(a.InitialPrice)
	a.HalfPercentOfInitialPrice = MarErm.PriceToString(initialPrice * 5 / 1000)
	a.FivePercentsOfInitialPrice = MarErm.PriceToString(initialPrice * 5 / 100)
	uadmin.Save(a)
}
