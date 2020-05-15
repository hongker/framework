package session

import (
	"github.com/gorilla/sessions"
	"github.com/hongker/framework/app"
	"github.com/rbcervilla/redisstore"
)

// NewCookieStore 返回普通的基于cookie机制的session
func NewCookieStore(key string) *sessions.CookieStore {
	return sessions.NewCookieStore([]byte(key))
}

// NewRedisStore 返回基于redis存储机制的session
func NewRedisStore() (*redisstore.RedisStore, error) {
	store, err := redisstore.NewRedisStore(app.Redis())
	if err != nil {
		return nil, err
	}

	return store, nil
}
