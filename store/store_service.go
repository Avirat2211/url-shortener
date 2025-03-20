package store

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type StorageService struct {
	redisClient *redis.Client
}

var (
	storeService = &StorageService{}
	ctx          = context.Background()
)

const cacheDuration = 6 * time.Hour

func InitializeStore() *StorageService {
	redisClient := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	pong, err := redisClient.Ping(ctx).Result()
	if err != nil {
		panic(fmt.Sprintf("Error init Redis: %v", err))
	}
	fmt.Printf("\n Redis started successfully: pong message = {%s} ", pong)
	storeService.redisClient = redisClient
	return storeService
}

func SaveUrlMapping(shortURL string, originalURL string, userId string) {
	err := storeService.redisClient.Set(ctx, shortURL, originalURL, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortURL: %s - originalURL: %s\n", err, shortURL, originalURL))
	}
}

func RetriveInitialUrl(shortURL string) string {
	result, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed retrieving url | Error : %v - shortURL: %s", err, shortURL))
	}
	return result
}
