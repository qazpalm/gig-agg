package apikeys

import (
	"sync"
)

var apiKeyArray = []string{
	"api_key_1",
}

type APIKeyManager struct {
	mu sync.Mutex
}

func NewAPIKeyManager() *APIKeyManager {
	return &APIKeyManager{}
}

func (akm APIKeyManager) IsValid(key string) bool {
	akm.mu.Lock()
	defer akm.mu.Unlock()

	for _, k := range apiKeyArray {
		if key == k {
			return true
		}
	}
	return false
}

