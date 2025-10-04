package redis

import (
	"github.com/go-redis/redis"
	"muthukamalan.io/shortnerurl/shortener"
)

type redisRepository struct {
	client *redis.Client
}

func newRedisClient(redisURL string) (*redis.Client, error) {
	if opts, err := redis.ParseURL(redisURL); err != nil {
		return nil, err
	} else {
		client := redis.NewClient(opts)
		_, err = client.Ping().Result()
		return client, err
	}
}

func NewRedisRepository(redisURL string) shortener.IRedirectRepository {

}

func (redisRepo *redisRepository) generate() {
}

func (redisRepo *redisRepository) Fetch() {

}
