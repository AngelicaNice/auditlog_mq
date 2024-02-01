package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/AngelicaNice/auditlog_mq/internal/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConnectionInfo struct {
	DbName   string
	URI      string
	Username string
	Password string
}

func InitDB(cfg *config.Config) *mongo.Database {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	opts := options.Client()
	opts.SetAuth(options.Credential{
		Username: cfg.Mongo.Username,
		Password: cfg.Mongo.Password,
	})
	opts.ApplyURI(cfg.Mongo.URI)

	dbClient, err := mongo.Connect(ctx, opts)
	if err != nil {
		log.Fatal(err)
	}

	if err := dbClient.Ping(context.Background(), nil); err != nil {
		log.Fatal(err)
	}

	fmt.Println("MONGODB CONNECTED")

	return dbClient.Database(cfg.Mongo.DbName)
}
