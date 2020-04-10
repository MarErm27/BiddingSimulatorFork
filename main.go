package main

import (
	"github.com/MarErm27/BiddingSimulator/models"
	"github.com/MarErm27/uadmin"
)

func main() {
	uadmin.Register(
		models.Auction{},
	)
	uadmin.Port = 443
	uadmin.RootURL = "/admin/"
	// setting := uadmin.Setting{}
	// uadmin.Update(&setting, "Value", strconv.Itoa(uadmin.Port), "code = ?", "uAdmin.Port")
	// uadmin.Update(&setting, "Value", uadmin.RootURL, "code = ?", "uAdmin.RootURL")
	uadmin.SiteName = "Auctions and analytics"
	uadmin.StartServer()
}
