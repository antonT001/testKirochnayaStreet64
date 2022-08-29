package log

import (
	"encoding/json"
	"fmt"
	"gettingLogs/internal/logger"
	"gettingLogs/internal/models"
	s "gettingLogs/internal/repository/log"
	"gettingLogs/internal/vo"
	"net"
	"net/http"
	"strings"
	"sync"

	"github.com/google/uuid"
)

const (
	VALIDATE_ERROR              = "validate error"
	JSON_PARSE_ERROR            = "error parse json"
	VALIDATE_UUID_ERROR         = "error validate uuid"
	VALIDATE_IP_ERROR           = "error validate ip"
	VALIDATE_USER_UUID_ERROR    = "error validate user uuid"
	VALIDATE_TIMESTAMP_ERROR    = "error validate timestamp"
	VALIDATE_URL_ERROR          = "error validate url"
	VALIDATE_DATAREQUEST_ERROR  = "error validate datarequest"
	VALIDATE_DATARESPONSE_ERROR = "error validate dataresponse"
)

type Queue struct {
	logService s.Log
	logger     logger.Logger
	bigData    []s.LogAdd
	iventCh    chan struct{}
	mu         *sync.Mutex
}

func New(logService s.Log, logger logger.Logger) *Queue {
	var mu sync.Mutex
	var bigData []s.LogAdd
	iventCh := make(chan struct{})
	return &Queue{
		logService: logService,
		logger:     logger,
		bigData:    bigData,
		iventCh:    iventCh,
		mu:         &mu,
	}
}

func (b *Queue) TakeFromTheQueue() {
	for {

		if len(b.bigData) == 0 {
			select {
			case <-b.iventCh:
			}
			continue
		}
		b.mu.Lock()
		data := b.bigData[0]
		b.mu.Unlock()

		err := b.logService.Add(&data)
		if err != nil {
			b.logger.Log(err) //проверить
		} else {
			b.mu.Lock()
			b.bigData = b.bigData[1:]
			b.mu.Unlock()
		}
	}
}

func (b *Queue) AddToQueue(dataByte []byte, r *http.Request) error {
	data, err := validateLog(dataByte, r)
	if err != nil {
		return err
	}

	b.mu.Lock()
	b.bigData = append(b.bigData, *data)
	b.mu.Unlock()
	b.iventCh <- struct{}{}
	return nil
}

func validateLog(data []byte, r *http.Request) (*s.LogAdd, error) {
	logAddIn := models.LogAddIn{}
	
	err := json.Unmarshal(data, &logAddIn)
	if err != nil {
		return nil, fmt.Errorf(JSON_PARSE_ERROR)
	}

	logAdd := s.LogAdd{}
	logAdd.EventsVo = make([]s.EventsVo, len(logAddIn.Events))
	
	uuId := uuid.New().String()
	uuidVo, err := vo.ExamineUuid(uuId)
	if err != nil {
		return nil, fmt.Errorf(VALIDATE_UUID_ERROR)
	}
	logAdd.Uuid = uuidVo

	ip, _ := getIP(r)
	ipVo, err := vo.ExamineIntIp(ip)
	if err != nil {
		return nil, fmt.Errorf(VALIDATE_IP_ERROR)
	}
	logAdd.Ip = ipVo

	userUuid, err := vo.ExamineUuid(logAddIn.UserUuid)
	if err != nil {
		return nil, fmt.Errorf(VALIDATE_USER_UUID_ERROR)
	}
	logAdd.UserUuid = userUuid

	timestamp, err := vo.ExamineIntData(logAddIn.Timestamp)
	if err != nil {
		return nil, fmt.Errorf(VALIDATE_TIMESTAMP_ERROR)
	}
	logAdd.Timestamp = timestamp

	for i, events := range logAddIn.Events {
		url, err := vo.ExamineVoUrl(events.Url)
		if err != nil {
			return nil, fmt.Errorf(VALIDATE_URL_ERROR)
		}
		logAdd.EventsVo[i].Url = url

		dataRequest, err := vo.ExaminePayload(events.DataRequest)
		if err != nil {
			return nil, fmt.Errorf(VALIDATE_DATAREQUEST_ERROR)
		}
		logAdd.EventsVo[i].DataRequest = dataRequest

		dataResponse, err := vo.ExaminePayload(events.DataResponse)
		if err != nil {
			return nil, fmt.Errorf(VALIDATE_DATARESPONSE_ERROR)
		}
		logAdd.EventsVo[i].DataResponse = dataResponse
	}

	return &logAdd, nil
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
