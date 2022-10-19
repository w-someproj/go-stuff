package main

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	todo "github.com/w-someproj/go-stuff/applications/todoapp"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/handler"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/repository"
	"github.com/w-someproj/go-stuff/applications/todoapp/pkg/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error on initializing config: %s", err.Error())
	}

	if err := godotenv.Load("applications/todoapp/.env"); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	defer db.Close()
	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)

	// for graceful shutdown (end all db queries which started)
	go func() {
		if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running http server: %s", err.Error())
		}
	}()

	log.Println("todoapp Started")

	// blocking main from ending
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Println("todoapp End")
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf(`Error on shutting down: %s`, err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("applications/todoapp/configs")
	viper.SetConfigName("config") // set name of file
	return viper.ReadInConfig()
}
