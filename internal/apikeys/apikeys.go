package apikeys

import (
	"sync"
)

const apiKeyArray = []string{
	"api_key_1",
}

type apiKeyManager struct {
	mu sync.Mutex
}

func NewAPIKeyManager() *apiKeyManager {
	return &APIKeyManager{}
}

func (akm APIKeyManager) Validate(key string) bool {
	akm.mu.Lock()
	defer akm.mu.Unlock()

	for _, k := range apiKeyArray {
		if key == k {
			return true
		}
	}
	return false
}

