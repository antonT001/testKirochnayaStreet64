package log

import (
	"gettingLogs/internal/logger"
	queue "gettingLogs/internal/queue/log"
	logService "gettingLogs/internal/service/log"
	"net/http"
)

type Log interface {
	Add(w http.ResponseWriter, r *http.Request)
}

type log struct {
	logService logService.Log
	logger     logger.Logger
	queue      *queue.Queue
}

func New(logService logService.Log, logger logger.Logger, queue *queue.Queue) Log {
	return &log{logService: logService, logger: logger, queue: queue}
}
