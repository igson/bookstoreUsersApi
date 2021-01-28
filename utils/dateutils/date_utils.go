package dateutils

import (
	"time"
)

const (
	apiDateLayout     = "2006-01-02T15:04:05Z"
	apiDatabaseLayout = "2006-01-02T15:04:05"
)

//GetNow retorna a data atual do servidor
func GetNow() time.Time {
	return time.Now().UTC()
}

//GetNowString retorna a data atual do servidor
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat retorna a data atual do banco de dados
func GetNowDBFormat() string {
	return GetNow().Format(apiDatabaseLayout)
}
