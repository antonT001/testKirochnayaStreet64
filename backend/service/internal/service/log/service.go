package log

import (
	"gettingLogs/internal/logger"
	logRepository "gettingLogs/internal/repository/log"
)

type Log interface {
	Add(logAddIn *LogAddIn) error
}

type log struct {
	logRepository logRepository.Log
	logger        logger.Logger
}

func New(logRepository logRepository.Log, logger logger.Logger) Log {
	return &log{logRepository: logRepository, logger: logger}
}
