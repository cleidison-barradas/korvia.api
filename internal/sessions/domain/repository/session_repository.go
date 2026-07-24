package repository

import (
	"context"

	"github.com/cleidison-barradas/korvia.api/internal/sessions/domain/model"
)

type SessionRepository interface {
	SetSession(ctx context.Context, session *model.Session) error
	GetSession(ctx context.Context, phoneNumber string) (*model.Session, error)
}