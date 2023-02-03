package httpserver

import (
	"net/http"
	"net/http/cookiejar"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

func Login(c *fiber.Ctx) error {
	type Payload struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	var p Payload

	if err := c.BodyParser(&p); err != nil {
		return err
	}

	jar, err := cookiejar.New(nil)
	if err != nil {
		return err
	}
	client := &http.Client{
		Jar: jar,
	}

	if _, err = client.PostForm("https://academic.ui.ac.id/main/Authentication/Index", url.Values{"u": {p.Username}, "p": {p.Password}}); err != nil {
		return err
	}

	resp, err := client.Get("https://academic.ui.ac.id/main/Authentication/ChangeRole")
	if err != nil {
		return err
	}

	for h := range resp.Header {
		c.Set(h, resp.Header.Get(h))
	}
	for _, cookie := range jar.Cookies(resp.Request.URL) {
		cookies := &fiber.Cookie{}
		cookies.Name = cookie.Name
		cookies.Value = cookie.Value
		cookies.Path = cookie.Path
		cookies.Domain = cookie.Domain
		cookies.MaxAge = cookie.MaxAge
		cookies.Expires = cookie.Expires
		cookies.Secure = cookie.Secure
		cookies.HTTPOnly = cookie.HttpOnly

		c.Cookie(cookies)
	}

	return c.Status(resp.StatusCode).Send(nil)
}
