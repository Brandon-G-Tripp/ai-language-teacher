package auth 

import "sync"

type tokenInvalidator struct {
   invalidTokens map[string]struct{} 
   mu sync.RWMutex
} 

var invalidator = tokenInvalidator{
    invalidTokens: make(map[string]struct{}),
} 

func (i *tokenInvalidator) InvalidateToken(token string) {
    i.mu.Lock()
    defer i.mu.Unlock()
    i.invalidTokens[token] = struct{}{}
} 

func (i *tokenInvalidator) IsTokenInvalid(token string) bool {
    i.mu.RLock()
    defer i.mu.RUnlock()
    _, ok := i.invalidTokens[token]
    return ok
} 
