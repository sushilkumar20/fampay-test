package common

import "fmt"

type ErrorRequest struct {
	Payload string `json:"reason"`
}

func (r *ErrorRequest) Error() string {
	return fmt.Sprintf("{\"payload\":\"%s\"}", r.Payload)
}
