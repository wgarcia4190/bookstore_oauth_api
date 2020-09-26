package app

import (
	"github.com/gin-gonic/gin"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/clients/cassandra"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/domain/access_token"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/http"
	"github.com/wgarcia4190/bookstore_oauth_api/internal/repositories/db"
)

func StartApplication() {
	router := createRouter()
	session, dbErr := cassandra.GetSession()

	if dbErr != nil {
		panic(dbErr)
	}
	defer session.Close()

	repository := db.NewRepository(session)
	atHandler := http.NewHandler(access_token.NewService(repository))

	router.GET("/oauth/access_token/:access_token_id", atHandler.GetById)
	router.POST("/oauth/access_token", atHandler.Create)

	_ = router.Run(":8080")
}

// createRouter creates the gin.Engine object.
func createRouter() *gin.Engine {
	return gin.Default()
}
