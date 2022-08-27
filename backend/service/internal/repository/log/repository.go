package log

import (
	"gettingLogs/internal/clients"
	"gettingLogs/internal/logger"
)

type Log interface {
	Add(valueStrings string, valueArgs []interface{}) error
}

type log struct {
	db     clients.DataBase
	logger logger.Logger
}

func New(db clients.DataBase, logger logger.Logger) Log {
	return &log{
		db:     db,
		logger: logger,
	}
}
