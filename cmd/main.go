package main

import (
	"net"
	"os"
	"os/signal"
	pbps "post-service/genproto/post_service" // protobuffer post service
	postService "post-service/internal/postService"
	"post-service/pkg/config"
	"post-service/pkg/db"
	"post-service/pkg/logger"
	post_repo "post-service/repo"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {

	cfg := config.NewConfig()

	db, logger :=db.NewDB(cfg), logger.New(cfg.GetString("app.log.level"), cfg.GetString("app.name"))

	postRepo := post_repo.NewPostRepo(db, logger)

	service := postService.NewPostService(postRepo, logger)

	listener, err := net.Listen("tcp", cfg.GetString("app.port"))
	if err != nil {
		panic(err)
	}

	s := grpc.NewServer()
	pbps.RegisterPostServiceServer(s, service)
	reflection.Register(s)

	shutDownCh := make(chan os.Signal,1)

	signal.Notify(shutDownCh, os.Interrupt)

	go func ()  {
		<- shutDownCh
		s.GracefulStop()
		logger.Info("shutting down gracefully")
	}()

	logger.Info("service has started it's job on port: " + cfg.GetString("app.port"))

	if err = s.Serve(listener); err != nil {
		panic(err)
	}

	logger.Info("server is shutted down")
}
