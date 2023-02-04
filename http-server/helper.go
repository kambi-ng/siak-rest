package httpserver

import (
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type SiakError struct {
	Status  int
	Message string
}

func (e *SiakError) Error() string {
	return fmt.Sprintf("SIAK error: %s", e.Message)
}

func MakeRequestor(c *fiber.Ctx) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if strings.Contains(req.URL.Path, "Authentication") {
				return &SiakError{Status: fiber.StatusUnauthorized, Message: "Session has expired"}
			}

			return nil
		},
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
