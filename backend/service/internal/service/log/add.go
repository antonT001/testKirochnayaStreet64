package log

import (
	s "gettingLogs/internal/repository/log"
)

type LogAddOut struct {
	Success bool    `json:"success"`
	Error   *string `json:"error,omitempty"`
}

func (s *log) Add(logAdd *s.LogAdd) error {
	return s.logRepository.Add(logAdd)
}
