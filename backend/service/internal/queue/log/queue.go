package log

import (
	"gettingLogs/internal/logger"
	logService "gettingLogs/internal/service/log"
	"sync"
)

type Queue struct {
	logService logService.Log
	logger     logger.Logger
	bigData    []logService.LogAddIn
	ch         chan logService.LogAddIn
	iventCh    chan struct{}
	mu         *sync.Mutex
}

func New(
	logService logService.Log,
	logger logger.Logger,
	bigData []logService.LogAddIn,
	ch chan logService.LogAddIn,
	iventCh chan struct{},
	mu *sync.Mutex) *Queue {
	return &Queue{
		logService: logService,
		logger:     logger,
		bigData:    bigData,
		ch:         ch,
		iventCh:    iventCh,
		mu:         mu,
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

func (b *Queue) AddToQueue(data *logService.LogAddIn) {
	b.mu.Lock()
	b.bigData = append(b.bigData, *data)
	b.mu.Unlock()
	b.iventCh <- struct{}{}
}
