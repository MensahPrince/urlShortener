package utils

import "sync"

var OTPStore = map[string]int{}
var mu sync.Mutex

type RATELIMITER struct {
	counts map[string]int
}

func (rl *RATELIMITER) Allow(ip string, limit int) bool {
	mu.Lock()
	defer mu.Unlock()

	if rl.counts == nil {
		rl.counts = make(map[string]int)
	}

	rl.counts[ip]++
	return rl.counts[ip] <= limit
}
