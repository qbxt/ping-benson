package main

import (
	"github.com/qbxt/status-benson/constants"
	"github.com/qbxt/status-benson/handlers"
	"github.com/qbxt/status-benson/fullInit"
	"github.com/qbxt/status-benson/logger"
	"github.com/gorilla/mux"
	"github.com/qbxt/status-benson/structures"
	"net/http"
)

func main() {

	logger.Init()

	err := fullInit.Init()
	if err != nil {
		logger.Error("could not fullInit Discord bot", err, nil)
		return
	}

	go processIncomingRequests(fullInit.IncomingPings)

	fullInit.Bot.AddHandler(handlers.OnMessage)
	fullInit.Bot.AddHandler(handlers.OnReady)
	go runDiscordBot()

	r := mux.NewRouter()
	r.HandleFunc("/", handlers.PingEndpoint)
	/* Example: benson.nmt.gg/ping?id=5&ping=54 */

	srv := &http.Server{
		Handler: r,
		Addr: "0.0.0.0:6969",
	}

	logger.Info("webserver is running", nil)
	logger.Fatal("server crashed", srv.ListenAndServe(), nil)
}

func runDiscordBot() {
	if err := fullInit.Bot.Open(); err != nil {
		logger.Error("could not open connection with discord", err, nil)
		return
	}
}

func processIncomingRequests(incomingPings chan structures.Ping) {
	for {
		newPing := <-incomingPings
		if checkID(newPing.ID) {
			fullInit.Pings[newPing.ID] = newPing
		}
	}
}

func checkID(ID string) bool {
	for e, _ := range constants.ApprovedIDs {
		if e == ID {
			return true
		}
	}
	return false
}