package main

import (
	"product"
	"product/pkg/handler"
	"product/pkg/repository"
	"product/pkg/service"
	"product/products"
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	err := godotenv.Load()
    if err != nil {
        log.Fatalf("err loading: %v", err)
    }
	conn, err := grpc.Dial(os.Getenv("SERVER_HOST"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect: %v", err)
	}
	defer conn.Close()
	// Вызов и обработка описанных методов

	client := products.NewProductsClient(conn)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("POSTGRES_HOST"),
		Port:     os.Getenv("POSTGRES_PORT"),
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		DBName:   os.Getenv("POSTGRES_DB"),
		SSLMode:  os.Getenv("POSTGRES_SSLMODE"),
	})
	if err != nil {
		log.Fatalf("error failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, client)
	handlers := handler.NewHandler(services)
	srv := new(product.Server)
	go func() {
		if err := srv.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
			log.Fatalf("error occured while running server: %s", err.Error())
		}
	}()

	log.Print("Product Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Print("Product Shutting Down")
	if err := srv.ShutDown(context.Background()); err != nil {
		log.Fatalf("error occured on server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		log.Fatalf("error occured on server while closing db: %s", err.Error())
	}
}
