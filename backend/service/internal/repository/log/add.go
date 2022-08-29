package log

import (
	"fmt"
	"gettingLogs/internal/vo"
	"strings"
)

type LogAdd struct {
	Uuid      vo.VoUuid
	Ip        vo.Ip
	UserUuid  vo.VoUuid
	Timestamp vo.IntData
	EventsVo  []EventsVo
}

type EventsVo struct {
	Url          vo.VoUrl
	DataRequest  vo.Payload
	DataResponse vo.Payload
}

func (r *log) Add(logAdd *LogAdd) error {
	valueStrings, valueArgs := bulkInsertPreparation(logAdd)
	return r.db.Exec(valueStrings, valueArgs...)
}

func bulkInsertPreparation(logAdd *LogAdd) (string, []interface{}) {
	valueStrings := make([]string, 0, len(logAdd.EventsVo))
	valueArgs := make([]interface{}, 0, len(logAdd.EventsVo)*7)
	for _, post := range logAdd.EventsVo {
		valueStrings = append(valueStrings, "(?, ?, ?, ?, ?, ?, ?)")
		valueArgs = append(valueArgs, logAdd.Uuid.Uuid())
		valueArgs = append(valueArgs, logAdd.Ip.Ip())
		valueArgs = append(valueArgs, logAdd.UserUuid.Uuid())
		valueArgs = append(valueArgs, logAdd.Timestamp.Data())
		valueArgs = append(valueArgs, post.Url.Url())
		valueArgs = append(valueArgs, post.DataRequest.Payload())
		valueArgs = append(valueArgs, post.DataResponse.Payload())
	}
	stmt := fmt.Sprintf("INSERT INTO logs (uuid, ip, user_uuid, timestamp, url, dataRequest, dataResponse) VALUE %s", strings.Join(valueStrings, ","))

	return stmt, valueArgs
}
