package bot

import (
	"tg-bot/internal/cache"
	"tg-bot/internal/domain"
	"time"
)

type botStorage struct {
	tokenCache      cache.Cache[string]
	userActionCache cache.Cache[domain.UserAction]
}

func newBotStorage(expiration time.Duration) *botStorage {
	return &botStorage{
		tokenCache:      cache.NewTTLCache[string](expiration),
		userActionCache: cache.NewTTLCache[domain.UserAction](expiration),
	}
}

func (m *botStorage) CheckToken(username string) (string, bool) {
	token, ok := m.tokenCache.Get(username)
	if !ok {
		return "", false
	}
	return token, true
}

func (m *botStorage) SetToken(username, token string) {
	m.tokenCache.Set(username, token)
}

func (m *botStorage) GetLatestAction(username string) (domain.UserAction, bool) {
	return m.userActionCache.Get(username)
}

func (m *botStorage) SetLatestAction(username string, action domain.UserAction) {
	m.userActionCache.Set(username, action)
}
