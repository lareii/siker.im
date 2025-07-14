package middleware

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v3"
)

type TurnstileMiddleware struct {
	Secret string
}

type turnstileResponse struct {
	Success     bool     `json:"success"`
	ChallengeTS string   `json:"challenge_ts"`
	Hostname    string   `json:"hostname"`
	ErrorCodes  []string `json:"error-codes"`
}

func NewTurnstileMiddleware(secret string) *TurnstileMiddleware {
	return &TurnstileMiddleware{Secret: secret}
}

func (t *TurnstileMiddleware) Verify() fiber.Handler {
	return func(c fiber.Ctx) error {
		if t.Secret == "" {
			return c.Next()
		}

		token := c.Get("Cf-Turnstile-Token")
		if token == "" {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Missing Turnstile token",
			})
		}

		ip := c.IP()
		ok, err := t.verifyToken(token, ip)
		if err != nil || !ok {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Turnstile verification failed",
			})
		}

		return c.Next()
	}
}

func (t *TurnstileMiddleware) verifyToken(token, ip string) (bool, error) {
	resp, err := http.PostForm("https://challenges.cloudflare.com/turnstile/v0/siteverify",
		url.Values{
			"secret":   {t.Secret},
			"response": {token},
			"remoteip": {ip},
		})
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result turnstileResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.Success, nil
}
