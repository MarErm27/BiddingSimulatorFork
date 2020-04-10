package main

import (
	"net/http"
	"strconv"

	"github.com/MarErm27/BiddingSimulatorFork/models"
	"github.com/MarErm27/BiddingSimulatorFork/views"
	"github.com/MarErm27/uadmin"
)

func showMainPage(w http.ResponseWriter, r *http.Request) {
	// data := TodoPageData{
	// 	PageTitle: "My TODO list",
	// 	// Todos:     Auctions,
	// }
	mainPageTempName := "templates/mainPage.html"
	uadmin.RenderHTML(w, r, mainPageTempName, nil)
}

func main() {
	uadmin.Register(
		models.Auction{},
		models.Participant{},
		models.AuctionTable{},
	)

	uadmin.RegisterInlines(
		models.Auction{},
		map[string]string{
			"AuctionTable": "AuctionID",
			// "offer":          "AuctionTableID",
		},
	)

	uadmin.RegisterInlines(
		models.Participant{},
		map[string]string{
			"AuctionTable": "ParticipantID",
			// "offer":          "AuctionTableID",
		},
	)
	uadmin.SiteName = "Auctions and analytics"
	uadmin.Port = 443
	uadmin.RootURL = "/admin/"
	setting := uadmin.Setting{}
	uadmin.Update(&setting, "Value", strconv.Itoa(uadmin.Port), "code = ?", "uAdmin.Port")
	uadmin.Update(&setting, "Value", uadmin.RootURL, "code = ?", "uAdmin.RootURL")
	// uadmin.Update(&setting, "Value", "default", "code = ?", "uAdmin.Theme")
	http.HandleFunc("/", http.HandlerFunc(showMainPage))
	http.HandleFunc("/http_handler/", views.HTTPHandler)
	uadmin.StartServer()
}
