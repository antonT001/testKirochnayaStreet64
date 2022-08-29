package vo

import (
	"fmt"
)

type Payload struct {
	value string
}

func ExaminePayload(value string) (Payload, error) {
	var (
		payload Payload
		err     error
	)

	payload.value = value

	if payload.value == "" {
		err = fmt.Errorf("payload error")
	}

	return payload, err
}

func (payload *Payload) Payload() string {
	return payload.value
}
