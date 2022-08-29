package log

import (
	"encoding/json"
	"gettingLogs/internal/helpers"
	s "gettingLogs/internal/service/log"
	"io/ioutil"
	"net/http"
)

const (
	SYSTEM_ERROR       = "system error"
	SERVICE_ERROR      = "service error"
	GETTING_IP_ERROR   = "getting ip error"
	ADD_TO_QUEUE_ERROR = "add to queue error"
)

func (h *log) Add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseAdd(w, s.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(SYSTEM_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	err = h.queue.AddToQueue(bodyBytes, r)
	if err != nil {
		responseAdd(w, s.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(ADD_TO_QUEUE_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	responseAdd(w, s.LogAddOut{
		Success: true,
	}, http.StatusOK)
}

func responseAdd(w http.ResponseWriter, out s.LogAddOut, statusCode int) {
	result, err := json.Marshal(out)
	if err != nil {
		w.Write([]byte(SYSTEM_ERROR + ": " + err.Error()))
		return
	}
	w.WriteHeader(statusCode)

	w.Write(result)
}
