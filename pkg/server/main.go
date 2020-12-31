package server

import (
	"github.com/fdelgato/go-gql-server/internal/handlers"
	log "github.com/fdelgato/go-gql-server/internal/logger"
	"github.com/fdelgato/go-gql-server/internal/orm"
	"github.com/fdelgato/go-gql-server/pkg/utils"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

var host, port, gqlPath, gqlPgPath string
var isPgEnabled bool

func init() {
	host = utils.MustGet("GQL_SERVER_HOST")
	port = utils.MustGet("GQL_SERVER_PORT")
	gqlPath = utils.MustGet("GQL_SERVER_GRAPHQL_PATH")
	gqlPgPath = utils.MustGet("GQL_SERVER_GRAPHQL_PLAYGROUND_PATH")
	isPgEnabled = utils.MustGetBool("GQL_SERVER_GRAPHQL_PLAYGROUND_ENABLED")
}

// Run spins up the server
func Run(orm *orm.ORM) {
	log.Info("GORM_CONNECTION_DSN: ", utils.MustGet("GORM_CONNECTION_DSN"))

	endpoint := "http://" + host + ":" + port

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		// AllowOriginFunc:  func(origin string) bool { return origin == "http://localhost:3000" },
		AllowOriginFunc:  func(origin string) bool { return true },
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Handlers
	// Simple keep-alive/ping handler
	r.GET("/heartbeat", handlers.Heartbeat())

	// GraphQL handlers
	// Playground handler
	if isPgEnabled {
		r.GET(gqlPgPath, handlers.PlaygroundHandler(gqlPath))
		log.Info("GraphQL Playground @ " + endpoint + gqlPgPath)
	}
	// Pass in the ORM instance to the GraphqlHandler
	r.POST(gqlPath, handlers.GraphqlHandler(orm))
	log.Info("GraphQL @ " + endpoint + gqlPath)

	// Run the server
	// Inform the user where the server is listening
	log.Info("Running @ " + endpoint)
	// Print out and exit(1) to the OS if the server cannot run
	log.Fatal(r.Run(host + ":" + port))
}


