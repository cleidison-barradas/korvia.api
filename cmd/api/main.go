package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/cleidison-barradas/korvia.api/internal/config"
	"github.com/cleidison-barradas/korvia.api/internal/database"
	"github.com/cleidison-barradas/korvia.api/internal/database/redis"
	"github.com/cleidison-barradas/korvia.api/internal/establishments"
	"github.com/cleidison-barradas/korvia.api/internal/messages"
	"github.com/cleidison-barradas/korvia.api/internal/services"
	"github.com/cleidison-barradas/korvia.api/internal/users"
	"github.com/cleidison-barradas/korvia.api/pkg/router"
	"github.com/cleidison-barradas/korvia.api/pkg/server"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	defer stop()

	cfg, err := config.Load()

	if err != nil {
		log.Fatal("Error on load config", err)
	}

	rdb := redis.NewClient(redis.Config{
		RedisHost: cfg.RedisHost,
		Password: cfg.RedisPassword,
		DB: cfg.RedisDB,
	})

	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatal("Error on start", err)
	}

	defer rdb.Close()

	pool, err := database.NewPool(ctx, cfg.DatabaseURl)

	if err != nil {
		log.Fatal("Error on start", err)
	}

	defer pool.Close()

	if err := database.RunMigrations(cfg.DatabaseURl); err != nil {
		log.Fatal("Error on run migrations", err)
	}

	handlers := []router.Register{
		users.NewModule(ctx, pool),
		messages.NewModule(ctx, rdb, pool),
		establishments.NewModule(ctx, pool),
		services.NewModule(ctx, pool),
	}

	srv, err := server.New(handlers...)

	if err != nil {
		log.Fatal("Error on register routes", err)
	}

	if err := srv.Listen(":" + cfg.ServerPort); err != nil {
		log.Fatal("Error on start server", err)
	}

	<-ctx.Done()
	stop()
	
	shutdowCtx, cancel := context.WithTimeout(context.Background(), 10*time.Microsecond)

	defer cancel()

	if err := srv.Shutdown(shutdowCtx); err != nil {
		log.Fatal("Error on shutdown server", err)
	}
}