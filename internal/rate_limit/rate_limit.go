package ratelimit

import (
	"sync"
	"time"
)

type TokenBucket struct {
	capacity   float64
	refillRate float64
	tokens     float64
	last       time.Time
	mu         sync.Mutex
}

func NewTokenBucket(capacity, refillRate float64) *TokenBucket {
	return &TokenBucket{
		capacity:   capacity,
		refillRate: refillRate,
		tokens:     capacity,
		last:       time.Now(),
	}
}

func (b *TokenBucket) Allow() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	now := time.Now()
	elapsed := now.Sub(b.last).Seconds()
	b.tokens += elapsed * b.refillRate
	if b.tokens > b.capacity {
		b.tokens = b.capacity
	}
	b.last = now

	if b.tokens >= 1 {
		b.tokens--
		return true
	}
	return false
}

type Limiter struct {
	mu         sync.Mutex
	buckets    map[string]*TokenBucket
	capacity   float64
	refillRate float64
}

func NewLimiter(capacity, refillRate float64) *Limiter {
	return &Limiter{
		buckets:    make(map[string]*TokenBucket),
		capacity:   capacity,
		refillRate: refillRate,
	}
}

func (l *Limiter) AllowIP(ip string) bool {
	l.mu.Lock()
	b, ok := l.buckets[ip]
	if !ok {
		b = NewTokenBucket(l.capacity, l.refillRate)
		l.buckets[ip] = b
	}
	l.mu.Unlock()

	return b.Allow()
}
