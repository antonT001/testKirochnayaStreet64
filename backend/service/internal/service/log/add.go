package log

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type LogAddIn struct {
	Uuid      string
	Ip        string
	UserUuid  string   `json:"user_uuid"`
	Timestamp uint32   `json:"timestamp"`
	Events    []Events `json:"events"`
}

type Events struct {
	Url          string `json:"url"`
	DataRequest  string `json:"dataRequest"`
	DataResponse string `json:"dataResponse"`
}

type LogAddOut struct {
	Success bool    `json:"success"`
	Error   *string `json:"error,omitempty"`
}

func (s *log) Add(logAddIn *LogAddIn) error {
	logAddIn.Uuid = uuid.New().String()
	valueStrings, valueArgs := bulkInsertPreparation(logAddIn)
	return s.logRepository.Add(valueStrings, valueArgs)
}

func bulkInsertPreparation(logAddIn *LogAddIn) (string, []interface{}) {
	valueStrings := make([]string, 0, len(logAddIn.Events))
	valueArgs := make([]interface{}, 0, len(logAddIn.Events)*7)
	for _, post := range logAddIn.Events {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, logAddIn.Uuid)
		valueArgs = append(valueArgs, logAddIn.Ip)
		valueArgs = append(valueArgs, logAddIn.UserUuid)
		valueArgs = append(valueArgs, logAddIn.Timestamp)
		valueArgs = append(valueArgs, post.Url)
		valueArgs = append(valueArgs, post.DataRequest)
		valueArgs = append(valueArgs, post.DataResponse)
	}
	stmt := fmt.Sprintf("INSERT INTO logs (uuid, ip, user_uuid, timestamp, url, dataRequest, dataResponse) VALUE %s", strings.Join(valueStrings, ","))

	return stmt, valueArgs
}
