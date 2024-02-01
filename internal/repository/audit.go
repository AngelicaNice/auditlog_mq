package repository

import (
	"context"

	audit "github.com/AngelicaNice/auditlog_mq/pkg/domain"

	"go.mongodb.org/mongo-driver/mongo"
)

type Audit struct {
	db *mongo.Collection
}

func NewAudit(db *mongo.Database, collection string) *Audit {
	return &Audit{
		db: db.Collection(collection),
	}
}

func (r *Audit) Insert(ctx context.Context, item audit.LogItem) error {
	_, err := r.db.InsertOne(ctx, item)

	return err
}
