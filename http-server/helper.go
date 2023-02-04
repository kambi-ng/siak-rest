package httpserver

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func MakeRequestor(c *fiber.Ctx) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}

	u, _ := url.Parse("https://academic.ui.ac.id")
	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "Mojavi",
			Value: c.Cookies("Mojavi"),
		},
		{
			Name:  "siakng_cc",
			Value: c.Cookies("siakng_cc"),
		},
	})

	return client, nil
}
