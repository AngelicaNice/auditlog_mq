package service

import (
	"context"

	audit "github.com/AngelicaNice/auditlog_mq/pkg/domain"
)

type AuditRepository interface {
	Insert(ctx context.Context, item audit.LogItem) error
}

type AuditService struct {
	repo AuditRepository
}

func NewAuditService(repo AuditRepository) *AuditService {
	return &AuditService{
		repo: repo,
	}
}

func (a *AuditService) Log(ctx context.Context, item audit.LogItem) error {
	return a.repo.Insert(ctx, item)
}
