package vo

import (
	"fmt"
	"time"
)

type IntData struct {
	value uint32
}

func ExamineIntData(value uint32) (IntData, error) {
	var (
		data IntData
		err  error
	)
	data.value = value
	if data.value < 1661714799 || data.value > uint32(time.Now().Unix()) {
		err = fmt.Errorf("data error")
	}
	return data, err
}

func (data *IntData) Data() uint32 {
	return data.value
}
