package routes

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

type request struct {
	URL         string `json:"url"`
	CustomShort string `json:"short"`
	Expiry      string `json:"expiry"`
}

type response struct {
	URL            string        `json:"url"`
	CustomShort    string        `json:"short"`
	Expiry         time.Duration `json:"expiry"`
	XRateRemaining int           `json:"rate_limit_remaining"`
	XRateLimitRest time.Duration `json:"rate_limit_rest"`
}

func ShortenURL(c *fiber.Ctx) error {
	var body := new(request)
}
