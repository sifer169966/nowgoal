package protocol

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"log"
// 	"net"
// 	"os"
// 	"os/signal"
// 	"nowgoal/configs"
// 	v1 "nowgoal/internal/api/v1"
// 	"nowgoal/internal/core/service"
// 	"nowgoal/internal/handlers"
// 	"nowgoal/internal/repositories"
// 	"nowgoal/pkg/converter"
// 	postgres "nowgoal/pkg/databases/postgres"
// 	"nowgoal/pkg/uidgen"

// 	"google.golang.org/grpc"
// )

// type grpcConfig struct {
// 	Env string
// }

// func ServeGRPC() error {
// 	ctx := context.Background()
// 	ctx, cancle := context.WithCancel(ctx)
// 	defer cancle()
// 	var cfg grpcConfig

// 	flag.StringVar(&cfg.Env, "env", "", "to pass environment via flag on command line")
// 	flag.Parse()

// 	configs.InitViper("./configs", cfg.Env)
// 	conn, err := postgres.ConnectPostgeSQL(
// 		configs.GetViper().DatabaseConfig.Postgres.Host,
// 		configs.GetViper().DatabaseConfig.Postgres.Port,
// 		configs.GetViper().DatabaseConfig.Postgres.Username,
// 		configs.GetViper().DatabaseConfig.Postgres.Password,
// 		configs.GetViper().DatabaseConfig.Postgres.DbName,
// 		configs.GetViper().Postgres.SSLMode,
// 	)
// 	if err != nil {
// 		return fmt.Errorf("Cannot connect to postgres database got error: '%s'", err)
// 	}

// 	postgresRepo := repositories.NewPostgres(conn.Postgres, converter.New())
// 	srv := service.New(postgresRepo, uidgen.New())
// 	handler := handlers.NewGRPCHandler(conn.Postgres, srv)

// 	opts := []grpc.ServerOption{}

// 	server := grpc.NewServer(opts...)
// 	v1.RegisterToteServiceServer(server, handler)
// 	listen, err := net.Listen("tcp", ":"+configs.GetViper().ServerConfig.GRPCPort)
// 	if err != nil {
// 		return fmt.Errorf("Cannot serve grpc server because: '%s'", err)
// 	}

// 	//Gracefull Shutdown ...
// 	c := make(chan os.Signal, 1)
// 	signal.Notify(c, os.Interrupt)
// 	go func() {
// 		for range c {
// 			//TODO: gracfull shutting down ...
// 			log.Println("Do something before shutdown ...")
// 			server.GracefulStop()
// 			<-ctx.Done()
// 		}
// 	}()
// 	return server.Serve(listen)
// }
