package did

import (
	"github.com/segmentio/ksuid"
)

func getSpecificID() string {
	id := ksuid.New()
	return id
}

moudule github.com/moonseokchoi-kr/did