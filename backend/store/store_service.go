package store

import (
	"context"
	"encoding/json"
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

type URLMapping struct {
	OriginalURL string `json:"original_url"`
	UserID      string `json:"user_id"`
}

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
	data := URLMapping{
		OriginalURL: originalURL,
		UserID:      userId,
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		panic(fmt.Sprintf("Failed to marshal URL mapping: %v", err))
	}
	err = storeService.redisClient.Set(ctx, shortURL, jsonData, cacheDuration).Err()
	if err != nil {
		panic(fmt.Sprintf("Failed saving key url | Error: %v - shortURL: %s - originalURL: %s\n", err, shortURL, originalURL))
	}
}

func RetriveInitialUrl(shortURL string) (string, string, error) {
	jsonData, err := storeService.redisClient.Get(ctx, shortURL).Result()
	if err != nil {
		return "", "", fmt.Errorf("failed retrieving url | Error: %v - shortURL: %s", err, shortURL)
	}
	var data URLMapping
	err = json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		return "", "", fmt.Errorf("failed unmarshaling url data | Error: %v - shortURL: %s", err, shortURL)
	}

	return data.OriginalURL, data.UserID, nil
}
