package log

import (
	"gettingLogs/internal/logger"
	logService "gettingLogs/internal/service/log"
	"net/http"
)

type Log interface {
	Add(w http.ResponseWriter, r *http.Request)
}

type log struct {
	logService logService.Log
	logger     logger.Logger
}

func New(logService logService.Log, logger logger.Logger) Log {
	return &log{logService: logService, logger: logger}
}
