package apikeys

import (
	"sync"
)

var apiKeyArray = []string{
	"api_key_1",
}

type ApiKeyManager struct {
	mu sync.Mutex
}

func NewAPIKeyManager() *ApiKeyManager {
	return &ApiKeyManager{}
}

func (akm ApiKeyManager) IsValid(key string) bool {
	akm.mu.Lock()
	defer akm.mu.Unlock()

	for _, k := range apiKeyArray {
		if key == k {
			return true
		}
	}
	return false
}

