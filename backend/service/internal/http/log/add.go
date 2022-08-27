package log

import (
	"encoding/json"
	"fmt"
	"gettingLogs/internal/helpers"
	logService "gettingLogs/internal/service/log"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

const (
	SYSTEM_ERROR     = "system error"
	SERVICE_ERROR    = "service error"
	JSON_PARSE_ERROR = "error parse json"
	GETTING_IP_ERROR = "getting ip error"
)

func (h *log) Add(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responseAdd(w, logService.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(SYSTEM_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	logAddIn, err := logUnmarshal(bodyBytes)
	if err != nil {
		responseAdd(w, logService.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(JSON_PARSE_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	logAddIn.Ip, err = getIP(r)
	if err != nil {
		responseAdd(w, logService.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(GETTING_IP_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	err = h.logService.Add(logAddIn)
	if err != nil {
		responseAdd(w, logService.LogAddOut{
			Success: false,
			Error:   helpers.StringPointer(SERVICE_ERROR + ": " + err.Error()),
		}, http.StatusInternalServerError)
		return
	}

	responseAdd(w, logService.LogAddOut{
		Success: true,
	}, http.StatusOK)
}

func responseAdd(w http.ResponseWriter, out logService.LogAddOut, statusCode int) {
	result, err := json.Marshal(out)
	if err != nil {
		w.Write([]byte(SYSTEM_ERROR + ": " + err.Error()))
		return
	}
	w.WriteHeader(statusCode)

	w.Write(result)
}

func logUnmarshal(bodyBytes []byte) (*logService.LogAddIn, error) {
	logAddIn := logService.LogAddIn{}

	err := json.Unmarshal(bodyBytes, &logAddIn)
	if err != nil {
		return nil, err
	}
	return &logAddIn, nil
}

func getIP(r *http.Request) (string, error) {
	//Get IP from the X-REAL-IP header
	ip := r.Header.Get("X-REAL-IP")
	netIP := net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	//Get IP from X-FORWARDED-FOR header
	ips := r.Header.Get("X-FORWARDED-FOR")
	splitIps := strings.Split(ips, ",")
	for _, ip := range splitIps {
		netIP := net.ParseIP(ip)
		if netIP != nil {
			return ip, nil
		}
	}

	//Get IP from RemoteAddr
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return "", err
	}
	netIP = net.ParseIP(ip)
	if netIP != nil {
		return ip, nil
	}

	return "", fmt.Errorf("no valid ip found")
}
