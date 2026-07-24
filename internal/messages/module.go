package messages

import (
	"context"
	"time"

	customerPostgresRepo "github.com/cleidison-barradas/korvia.api/internal/customers/infrastructure/persistence"
	"github.com/cleidison-barradas/korvia.api/internal/establishments/infrastructure/persistence"
	"github.com/cleidison-barradas/korvia.api/internal/messages/application/usecases"
	"github.com/cleidison-barradas/korvia.api/internal/messages/infraestructure/http"
	whatsappClient "github.com/cleidison-barradas/korvia.api/internal/messages/infraestructure/whatsapp"
	servicePostgresRepo "github.com/cleidison-barradas/korvia.api/internal/services/infraestructure/persistence"
	sessionRedisRepo "github.com/cleidison-barradas/korvia.api/internal/sessions/infrastructure/persistence"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/redis/go-redis/v9"
)

func NewModule(ctx context.Context, redis *redis.Client, pool *pgxpool.Pool) *http.WebhookHandler {

	sessionRedisRepo := sessionRedisRepo.NewSessionRedisRepository(&sessionRedisRepo.SessionRedisRepository{
		Redis: redis,
		TTL: 30*time.Minute,
	})

	whatsappSender := whatsappClient.NewClient("", "", "")
	
	customerPostgresRepo := customerPostgresRepo.NewCustomerPostgresRepository(ctx, pool)
	establishmentPostgresRepo := persistence.NewEstablishmentPostgresRepository(ctx, pool)
	servicePostgresRepo := servicePostgresRepo.NewServicesPostgresRepository(ctx, pool)
	
	uc := usecases.NewIncomingMessageUseCase(&usecases.IncomingMessageUseCase{
		Ctx: ctx,
		WhatsappSender: whatsappSender,
		SessionRepo: sessionRedisRepo,
		ServiceRepo: servicePostgresRepo,
		CustomerRepo: customerPostgresRepo,
		EstablishmentRepo: establishmentPostgresRepo,
	})

	return http.NewWebhookHandler(uc)
}