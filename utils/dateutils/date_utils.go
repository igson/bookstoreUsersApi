package dateutils

import (
	"time"
)

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
)

//GetNow retorna a data atual do servidor
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString retorna a data atual do servidor
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}
