package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func InitRedis() {
	// env
	if err := godotenv.Load(); err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}

	// connect to redis
	password := os.Getenv("REDIS_PASSWORD")
	uri := os.Getenv("REDIS_URI")
	
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     uri,
		Password: password,
	})

    _, err := RedisClient.Ping(context.Background()).Result()
    if err != nil {
        log.Fatalf("Ошибка при подключении к Redis: %v", err)
    }

    log.Println("Соединение с Redis установлено успешно.")
}

func SetData(key string, value interface{}, ttl time.Duration) error {
	ctx := context.Background()
	err := RedisClient.Set(ctx, key, value, ttl).Err()
	if err != nil {
		return err
	}
	log.Printf("Данные успешно записаны в Redis. Ключ: %s\n", key)
	return nil
}

func GetData(key string) (string, error) {
	ctx := context.Background()
	val, err := RedisClient.Get(ctx, key).Result()
	if err == redis.Nil {
		log.Printf("Данные не найдены для ключа: %s\n", key)
		return "", nil
	} else if err != nil {
		return "", err
	}
	log.Printf("Данные получены из Redis. Ключ: %s, Значение: %s\n", key, val)
	return val, nil
}

func Close() {
    if err := RedisClient.Close(); err != nil {
        log.Printf("Ошибка при закрытии соединения с Redis: %v", err)
    } else {
        log.Println("Соединение с Redis закрыто.")
    }
}