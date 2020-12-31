package main

import (
	log "github.com/fdelgato/go-gql-server/internal/logger"
	"github.com/fdelgato/go-gql-server/internal/orm"
	"github.com/fdelgato/go-gql-server/pkg/server"
)

func main() {
	// Create a new ORM instance to send it to our server
	orm, err := orm.Factory()
	if err != nil {
		log.Panic(err)
	}
	// Send: ORM instance
	server.Run(orm)
}

