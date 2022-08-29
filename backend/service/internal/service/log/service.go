package log

import (
	"gettingLogs/internal/logger"
	s "gettingLogs/internal/repository/log"
)

type Log interface {
	Add(logAdd *s.LogAdd) error
}

type log struct {
	logRepository s.Log
	logger        logger.Logger
}

func New(logRepository s.Log, logger logger.Logger) Log {
	return &log{logRepository: logRepository, logger: logger}
}
