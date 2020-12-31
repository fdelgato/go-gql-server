package utils

import (
	log "github.com/fdelgato/go-gql-server/internal/logger"
	"os"
	"strconv"
)

// MustGet will return the env or panic if it is not present
func MustGet(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.MissingArg(k)
		log.Panic("ENV missing, key: " + k)
	}
	return v
}

// MustGetBool will return the env as boolean or panic if it is not present
func MustGetBool(k string) bool {
	v := os.Getenv(k)
	if v == "" {
		log.Panic("ENV missing, key: " + k)
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		log.MissingArg(k)
		log.Panic("ENV err: [" + k + "]\n" + err.Error())
	}
	return b
}

