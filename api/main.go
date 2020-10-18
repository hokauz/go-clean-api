package api

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hokauz/go-clean-api/api/envs"
	handlerSample "github.com/hokauz/go-clean-api/api/handlers/sample"
	"github.com/hokauz/go-clean-api/api/pkg/db"
	"github.com/hokauz/go-clean-api/api/pkg/middleware"
	"github.com/hokauz/go-clean-api/api/pkg/router"
	coreSample "github.com/hokauz/go-clean-api/core/sample"
	"go.mongodb.org/mongo-driver/mongo"
)

// Start -
func Start() {
	// Define contexto e configurações de banco
	ctx := context.Background()
	set := envs.GetInfo()

	// Inicia banco de dados
	conn, err := db.Connect(ctx, set.DB)
	if err != nil {
		log.Fatal(err.Error())
	}
	defer conn.Client().Disconnect(ctx)

	// Inicia rota defaul da api
	rout := router.Setup()

	// Aplica cors e pre-flight
	rout.Use(middleware.Cors())

	// Adiciona rotas publicas e privadas
	public := rout.Group(set.PathAPI)
	private := rout.Group(set.PathAPI)
	private.Use(middleware.AuthRequired())

	// Prepare Sample Handler
	prepareSample(public, conn.Collection("sample"), "sample")

	// Start api
	rout.Run(set.Router.Port)
	fmt.Println("Start API")
}

func prepareSample(group *gin.RouterGroup, coll *mongo.Collection, name string) {
	sampleRepo := coreSample.NewMongoRespository(context.Background(), coll)
	sampleService := coreSample.NewService(sampleRepo)
	handlerSample.Start(group, *sampleService, name)
}
