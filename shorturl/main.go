package main

import (
	"flag"
	"fmt"
	"net"
	"shorturl/pkg/config"
	"shorturl/pkg/db/mysql"
	"shorturl/pkg/db/redis"
	"shorturl/pkg/log"

	proto "shorturl/proc"
	"shorturl/shorturl-server/server"

	"google.golang.org/grpc"
)

var (
	configFile = flag.String("config", "shorturl.config.yaml", "")
)

func main() {
	flag.Parse()
	// 初始化配置文件
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

	//初始化mysql
	mysql.InitMysql(cnf)

	//初始化redis
	redis.InitRedisPool(cnf)

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", cnf.Server.IP, cnf.Server.Port))
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	service := server.NewService(cnf, logger)

	proto.RegisterShortUrlServer(s, service)
	if err = s.Serve(lis); err != nil {
		log.Fatal(err)
	}

}
