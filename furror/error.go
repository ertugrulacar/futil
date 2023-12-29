package furror

import (
	"fmt"
	"time"
)

type furror struct {
	Code         int    `json:"code"`
	DetailedCode int    `json:"detailed_code,omitempty"`
	Message      string `json:"message"`
	Time         int64  `json:"time"`
}

func (e *furror) Error() string {
	return fmt.Sprint(*e)
}

func New(message string, statusCode int) *furror {
	return &furror{
		Code:    statusCode,
		Message: message,
		Time:    time.Now().UTC().Unix(),
	}
}

func (e *furror) SetDetailedCode(detailedCode int) *furror {
	e.DetailedCode = detailedCode
	return e
}
