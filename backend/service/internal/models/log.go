package models

type LogAddIn struct {
	UserUuid  string   `json:"user_uuid"`
	Timestamp uint32   `json:"timestamp"`
	Events    []Events `json:"events"`
}

type Events struct {
	Url          string `json:"url"`
	DataRequest  string `json:"dataRequest"`
	DataResponse string `json:"dataResponse"`
}