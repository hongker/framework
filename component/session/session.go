package session

import (
	"github.com/go-redis/redis"
	"github.com/gorilla/sessions"
	"github.com/rbcervilla/redisstore"
)

// NewCookieStore 返回普通的基于cookie机制的session
func NewCookieStore(key string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(key))
}

// NewRedisStoreWithConn 指定连接获取session store
func NewRedisStoreWithConn(client redis.UniversalClient) (*redisstore.RedisStore, error) {
	store, err := redisstore.NewRedisStore(client)
	if err != nil {
		return nil, err
	}

	return store, nil
}