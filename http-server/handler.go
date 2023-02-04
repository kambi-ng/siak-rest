package httpserver

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kambi-ng/siak-rest/siaklib"
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

	resp, err := client.PostForm("https://academic.ui.ac.id/main/Authentication/Index", url.Values{"u": {p.Username}, "p": {p.Password}})
	if err != nil {
		return err
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if strings.Contains(string(b), "Login Failed") {
		return errors.New("authentication failed")
	}

	resp, err = client.Get("https://academic.ui.ac.id/main/Authentication/ChangeRole")
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

	return c.Status(resp.StatusCode).JSON(Response[any]{
		Status:  200,
		Message: "Authentication success. Please use given cookie for next requests.",
		Data:    nil,
	})
}

func Home(c *fiber.Ctx) error {
	req, err := MakeRequestor(c)
	if err != nil {
		return err
	}

	resp, err := req.Get("https://academic.ui.ac.id/main/Welcome/")
	if err != nil {
		return err
	}

	data, err := siaklib.ParseWelcomePage(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(data)
}
