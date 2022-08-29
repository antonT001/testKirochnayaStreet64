package vo

import (
	"fmt"

	"github.com/google/uuid"
)

type VoUuid struct {
	value string
}

func ExamineUuid(value string) (VoUuid, error) {
	var (
		voUuid VoUuid
		err    error
	)

	voUuid.value = value

	_, err = uuid.Parse(voUuid.value)
	if err != nil {
		err = fmt.Errorf("uuid error")
	}
	return voUuid, err
}

func (voUuid *VoUuid) Uuid() string {
	return voUuid.value
}
