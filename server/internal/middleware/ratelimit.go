package middleware

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/lareii/siker.im/internal/config"
	"github.com/lareii/siker.im/internal/database"

	"github.com/gofiber/fiber/v3"
	"go.uber.org/zap"
)

type RateLimiter struct {
	redis  *database.Redis
	config *config.RateLimitConfig
	logger *zap.Logger
}

func NewRateLimiter(redis *database.Redis, config *config.RateLimitConfig, logger *zap.Logger) *RateLimiter {
	return &RateLimiter{
		redis:  redis,
		config: config,
		logger: logger,
	}
}

// func (rl *RateLimiter) shouldRateLimit(c fiber.Ctx) bool {
// 	path := c.Path()
// 	method := c.Method()

// 	if method == "POST" && strings.HasPrefix(path, "/urls") {
// 		return true
// 	}

// 	return false
// }

func (rl *RateLimiter) Middleware() fiber.Handler {
	return func(c fiber.Ctx) error {
		if !rl.config.Enabled {
			return c.Next()
		}

		// if !rl.shouldRateLimit(c) {
		// 	return c.Next()
		// }

		clientIP := c.IP()
		key := fmt.Sprintf("rate_limit:%s", clientIP)

		ctx := context.Background()

		blockKey := fmt.Sprintf("blocked:%s", clientIP)
		blocked, err := rl.redis.Get(ctx, blockKey)
		if err == nil && blocked == "1" {
			rl.logger.Warn("Request blocked due to rate limit",
				zap.String("ip", clientIP),
				zap.String("path", c.Path()),
			)
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Rate limit exceeded. Please try again later.",
				"retry_after": int(rl.config.BlockTime.Seconds()),
			})
		}

		count, err := rl.redis.Incr(ctx, key)
		if err != nil {
			rl.logger.Error("Failed to increment rate limit counter",
				zap.Error(err),
				zap.String("ip", clientIP),
			)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Internal server error",
			})
		}

		if count == 1 {
			if err := rl.redis.Expire(ctx, key, rl.config.Window); err != nil {
				rl.logger.Error("Failed to set rate limit expiration",
					zap.Error(err),
					zap.String("ip", clientIP),
				)
			}
		}

		if count > int64(rl.config.Requests) {
			if err := rl.redis.Set(ctx, blockKey, "1", rl.config.BlockTime); err != nil {
				rl.logger.Error("Failed to block IP",
					zap.Error(err),
					zap.String("ip", clientIP),
				)
			}

			rl.logger.Warn("Rate limit exceeded",
				zap.String("ip", clientIP),
				zap.String("path", c.Path()),
				zap.Int64("count", count),
				zap.Int("limit", rl.config.Requests),
			)

			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"error":       "Rate limit exceeded. IP blocked temporarily.",
				"retry_after": int(rl.config.BlockTime.Seconds()),
			})
		}

		// remaining := int64(rl.config.Requests) - count
		// if remaining < 0 {
		// 	remaining = 0
		// }
		remaining := max(int64(rl.config.Requests)-count, 0)

		c.Set("X-RateLimit-Limit", strconv.Itoa(rl.config.Requests))
		c.Set("X-RateLimit-Remaining", strconv.FormatInt(remaining, 10))
		c.Set("X-RateLimit-Reset", strconv.FormatInt(time.Now().Add(rl.config.Window).Unix(), 10))

		return c.Next()
	}
}

// func (rl *RateLimiter) GetRateLimitStatus(ctx context.Context, ip string) (requests int64, blocked bool, err error) {
// 	key := fmt.Sprintf("rate_limit:%s", ip)
// 	blockKey := fmt.Sprintf("blocked:%s", ip)

// 	countStr, err := rl.redis.Get(ctx, key)
// 	if err != nil {
// 		requests = 0
// 	} else {
// 		requests, _ = strconv.ParseInt(countStr, 10, 64)
// 	}

// 	blockedStr, err := rl.redis.Get(ctx, blockKey)
// 	blocked = (err == nil && blockedStr == "1")

// 	return requests, blocked, nil
// }

// func (rl *RateLimiter) UnblockIP(ctx context.Context, ip string) error {
// 	blockKey := fmt.Sprintf("blocked:%s", ip)
// 	rateLimitKey := fmt.Sprintf("rate_limit:%s", ip)

// 	return rl.redis.Del(ctx, blockKey, rateLimitKey)
// }
