package vo

import (
	"fmt"
	"net/url"
)

type VoUrl struct {
	value string
}

func ExamineVoUrl(value string) (VoUrl, error) {
	var (
		voUrl VoUrl
		err   error
	)
	voUrl.value = value

	_, err = url.Parse(voUrl.value)
	if err != nil {
		err = fmt.Errorf("url error")
	}
	return voUrl, err
}

func (voUrl *VoUrl) Url() string {
	return voUrl.value
}
