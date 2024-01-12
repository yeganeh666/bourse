package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"mabna/services/trades/configs"
	"mabna/services/trades/docs"
	"mabna/services/trades/handlers"
	repo "mabna/services/trades/repository"
	"mabna/services/trades/services"
	"mabna/shared/configext"
	"mabna/shared/netext"
	"mabna/shared/service"
	"net"
	"net/http"
)

// @BasePath /api
// mabna service godoc
func main() {
	conf := new(configs.Configs)
	loadedConfigs, err := configext.LoadConfigs("", configs.DefaultConfig, true, conf)
	if err != nil {
		log.Fatal(err)
	}
	conf = loadedConfigs.(*configs.Configs)
	logger := log.New()
	logger.SetReportCaller(true)

	repository, err := repo.NewRepository(logger, conf)
	if err != nil {
		panic("Failed to create repository")
	}

	reportsService := services.NewReportService(logger,
		repo.NewReportsRepository(repository))

	handler := handlers.NewReportsHandler(
		logger, fmt.Sprintf(":%v", conf.GRPCPort))
	initSwagger(conf)

	service.Serve(netext.Port(conf.GRPCPort), func(lst net.Listener) error {
		server := grpc.NewServer()
		reportsService.RegisterService(server)
		return server.Serve(lst)
	})
	service.Serve(netext.Port(conf.HttpPort), func(lst net.Listener) error {
		router := handlers.Register(handler)
		return http.Serve(lst, router)
	})

	service.Start(conf.ServiceName, conf.Version)

}

func initSwagger(conf *configs.Configs) {
	fmt.Println(conf.HttpHost, conf.HttpPort)
	docs.SwaggerInfo.Title = "bourse"
	docs.SwaggerInfo.Description = "Bourse Service : As a dedicated microservice, the project aims to deliver a lightweight and efficient solution for retrieving the latest trades. By incorporating Protobuf for serialization, Swagger for documentation, and Docker Compose for containerization, the microservice is structured to support streamlined development, deployment, and ongoing maintenance, aligning with best practices in microservices architecture.\n"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api"
	docs.SwaggerInfo.Host = fmt.Sprintf("%v:%v", conf.HttpHost, conf.HttpPort)
	docs.SwaggerInfo.Schemes = []string{"http"}

}
