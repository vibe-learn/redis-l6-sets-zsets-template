// Package main is the redis lesson `l3_sets_and_zsets` homework scaffold for Vibe Learn.
//
// Задача: leaderboard на ZSET: PUT /score, GET /top?n, GET /rank/:user.
// Реализуй функции ниже — сигнатуры и тестовая поверхность фиксированы;
// CI (.github/workflows/ci.yml) гоняет `go vet` и `go test ./...`.
// Подробности и критерии приёмки — в README.md.
//
// Клиент: github.com/redis/go-redis/v9 (поддерживает Cluster/Sentinel/Pipeline).
package main

import (
	"bufio"
	"context"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/redis/go-redis/v9"
)

// Histogram — сборщик латентностей для перцентилей (TODO: замени на HDR при желании).
type Histogram struct{ Samples []time.Duration }

// Profile — пример доменной структуры для hash/string бэкендов.
type Profile struct {
	Name   string
	Email  string
	Visits int64
}

// ----- config -----

// envOr returns the env var for `key` if set, else `fallback`.
func envOr(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}

// Addr — адрес Redis. Дефолт совпадает с docker-compose.yml.
func Addr() string {
	return envOr("REDIS_ADDR", "localhost:6379")
}

// NewClient собирает *redis.Client из env. Override REDIS_ADDR в тестах.
func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{Addr: Addr()})
}

// ----- TODO #1: AddScore -----
//
// ZINCRBY — атомарно прибавить delta к очкам user (concurrent-safe)
func AddScore(ctx context.Context, rdb *redis.Client, board, user string, delta float64) (float64, error) {
	// TODO: implement
	panic("AddScore: not implemented")
}

// ----- TODO #2: TopN -----
//
// ZREVRANGE WITHSCORES — вернуть топ-N игроков по убыванию
func TopN(ctx context.Context, rdb *redis.Client, board string, n int64) ([]redis.Z, error) {
	// TODO: implement
	panic("TopN: not implemented")
}

// ----- TODO #3: Rank -----
//
// ZREVRANK — позиция игрока (0-based) в рейтинге
func Rank(ctx context.Context, rdb *redis.Client, board, user string) (int64, error) {
	// TODO: implement
	panic("Rank: not implemented")
}

// _refs keeps imports live while the TODO bodies are unimplemented stubs.
// Удали эту функцию, когда реализуешь TODO выше.
var _refs = []any{
	(*bufio.Reader)(nil),
	(io.Writer)(nil),
	(http.Handler)(nil),
	Histogram{},
	Profile{},
	time.Second,
}

// ----- main entry -----

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	log.Printf("Vibe Learn — redis lesson %s scaffold up", "l3_sets_and_zsets")
	log.Printf("redis addr: %s", Addr())
	log.Printf("Реализуй TODO-функции, затем `go test ./...`. README.md содержит задачу.")

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	rdb := NewClient()
	defer rdb.Close()

	// Graceful shutdown so `go run .` is interactive — Ctrl-C exits cleanly.
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		log.Printf("shutdown signal received")
		cancel()
	}()
	<-ctx.Done()
}
