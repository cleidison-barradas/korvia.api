package persistence

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/cleidison-barradas/korvia.api/internal/sessions/domain/model"
	"github.com/redis/go-redis/v9"
)

type SessionRedisRepository struct {
	Redis *redis.Client
	TTL time.Duration
}

func NewSessionRedisRepository(p *SessionRedisRepository) *SessionRedisRepository {
	return &SessionRedisRepository{
		TTL: p.TTL,
		Redis: p.Redis,
	}
}

func (r *SessionRedisRepository) key(phoneNumber string) string {
	return fmt.Sprintf("session:%s", phoneNumber)
}

func (r *SessionRedisRepository) GetSession(ctx context.Context, phoneNumber string) (*model.Session, error) {
	raw, err := r.Redis.Get(ctx, r.key(phoneNumber)).Bytes()

	if err == redis.Nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var session model.Session

	if err := json.Unmarshal(raw, &session); err != nil {
		return nil, err
	}

	return &session, nil
}

func (r *SessionRedisRepository) SetSession(ctx context.Context, session *model.Session) error {
	raw, err := json.Marshal(session)

	if err != nil {
    return err
	}

	if err := r.Redis.Set(ctx, r.key(session.PhoneNumber), raw, r.TTL).Err(); err != nil {
		return err
	}
	
	return nil
}