package handlers

import (
	"github.com/qbxt/status-benson/constants"
	"github.com/qbxt/status-benson/fullInit"
	"github.com/qbxt/status-benson/logger"
	"github.com/qbxt/status-benson/structures"
	"github.com/sirupsen/logrus"
	"net/http"
	"time"
)

func PingEndpoint(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		vars := r.URL.Query()
		id, ok1 := vars["id"]
		value, ok2 := vars["ping"]
		ok3 := false
		if ok1 {
			ok3 = checkID(id[0])
		}
		logger.Debug("pingcheck", logrus.Fields{"ok": ok1, "ok2": ok2, "ok3": ok3})
		if ok2 && ok3 {
			newPing := structures.Ping{
				LastSeen: time.Now().UTC().Unix(),
				ID: id[0],
				Value: value[0],
			}
			logger.Info("got a new ping", logrus.Fields{"pingStruct": newPing})
			fullInit.IncomingPings <- newPing
			w.WriteHeader(http.StatusOK)
			return
		} else {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
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