package protocol

import (
	"flag"
	"fmt"
	"log"
	"nowgoal/configs"
	"nowgoal/internal/core/service"
	"nowgoal/internal/handlers"
	"nowgoal/internal/repositories"
	"nowgoal/pkg/converter"
	postrges "nowgoal/pkg/databases/postgres"
	"nowgoal/pkg/genkey"
	"nowgoal/pkg/middlewares"
	"nowgoal/pkg/uidgen"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
)

type config struct {
	Env string
}

func ServeHTTP() error {
	app := fiber.New()

	var cfg config
	flag.StringVar(&cfg.Env, "env", "", "the environment to use")
	flag.Parse()
	configs.InitViper("./configs", cfg.Env)

	dbConn, err := postrges.ConnectPostgeSQL(
		configs.GetViper().DatabaseConfig.Postgres.Host,
		configs.GetViper().DatabaseConfig.Postgres.Port,
		configs.GetViper().DatabaseConfig.Postgres.Username,
		configs.GetViper().DatabaseConfig.Postgres.Password,
		configs.GetViper().DatabaseConfig.Postgres.DbName,
	)
	if err != nil {
		return fmt.Errorf("Cannot connect to the postgres database: '%s'", err)
	}
	postgresRepo := repositories.NewPostgres(dbConn.Postgres, nil)
	srv := service.New(postgresRepo, uidgen.New(), converter.New())
	hdl := handlers.New(srv)
	publicKey := genkey.GenPublicKey(configs.GetViper().Key.Public)
	api := app.Group("/v1", middlewares.AuthMiddleware(publicKey))
	{
		statAPI := api.Group("/stats")
		statAPI.Post("", hdl.GetStatsPattern1)
	}
	{
		fileAPI := api.Group("/files")
		fileAPI.Post("", hdl.ReadStat)
	}
	port := configs.GetViper().ServerConfig.HTTPPort
	err = app.Listen(":" + port)
	if err != nil {
		return err
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Gracefull shut down ...")
			//TODO: close database or any connection before server has gone ...
			app.Shutdown()
		}
	}()
	fmt.Println("Listerning on port: ", port)
	return nil

}
