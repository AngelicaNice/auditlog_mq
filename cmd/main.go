package main

import (
	"github.com/AngelicaNice/auditlog_mq/internal/config"
	"github.com/AngelicaNice/auditlog_mq/internal/repository"
	"github.com/AngelicaNice/auditlog_mq/internal/server"
	"github.com/AngelicaNice/auditlog_mq/internal/service"
	"github.com/AngelicaNice/auditlog_mq/internal/transport/mq"
	"github.com/AngelicaNice/auditlog_mq/pkg/database"
	log "github.com/sirupsen/logrus"
)

const (
	CONFIG_DIR  = "configs"
	CONFIG_FILE = "main"
)

func main() {
	cfg, err := config.NewConfig(CONFIG_DIR, CONFIG_FILE)
	if err != nil {
		log.WithField("config", "wrong config params").Fatal(err)
	}

	//ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	//defer cancel()

	db := database.InitDB(cfg)
	if db == nil {
		log.WithField("db creating", "failed").Fatal(err)
	}

	//defer db.Client().Disconnect(ctx)

	amqpConn, err := mq.CreateMQConnection(cfg.MQ.URL)
	if err != nil {
		log.WithField("rabbitmq", "failed to connect").Fatal(err)
	}

	defer amqpConn.Close()

	logsRepo := repository.NewAudit(db, cfg.MQ.Name)
	auditService := service.NewAuditService(logsRepo)
	auditServer := server.NewAuditServer(amqpConn, cfg.MQ.Name, auditService)
	auditServer.Run()
}
