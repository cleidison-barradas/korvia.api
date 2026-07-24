package service

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type MessagingService struct {
	Ctx context.Context
	RedisClient *redis.Client
}

func NewMessagingService(redisClient *redis.Client, ctx context.Context) *MessagingService {
	return &MessagingService{
		Ctx: ctx,
		RedisClient: redisClient,
	}
}

func (s *MessagingService) GetSession(chatID string) (sessionID string, err error) {

	return "", nil
}