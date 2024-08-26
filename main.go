package main

import (
	"flag"
	"fmt"
	"mediahub/controller"
	"mediahub/middleware"
	"mediahub/pkg/config"
	"mediahub/pkg/log"
	"mediahub/pkg/storage/cos"
	"mediahub/routers"

	"github.com/gin-gonic/gin"
)

var configFile = flag.String("config", "dev.config.yaml", "")

func main() {
	flag.Parse()
	config.InitConfig(*configFile)
	cnf := config.GetConfig()
	fmt.Printf("%+v\n", cnf)

	log.SetLevel(cnf.Log.Level)
	log.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	log.SetPrintCaller(true)

	logger := log.NewLogger()
	logger.SetOutput(log.GetRotateWriter(cnf.Log.LogPath))
	logger.SetLevel(cnf.Log.Level)
	logger.SetPrintCaller(true)

	sf := cos.NewCosStorageFactory(cnf.Cos.BucketUrl, cnf.Cos.SecretId, cnf.Cos.SecretKey, "")
	controller := controller.NewController(sf, logger, cnf)

	gin.SetMode(cnf.Http.Mode)
	r := gin.Default()
	api := r.Group("/api")
	api.Use(middleware.Cors())
	routers.InitRouters(api, controller)
	r.Run(fmt.Sprintf("%s:%d", cnf.Http.IP, cnf.Http.Port))

	fmt.Printf("%+v\n", sf)
}
