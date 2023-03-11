package httpserver

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/kambi-ng/siak-rest/siaklib"
)

type CookieData struct {
	SiakNGCC string `json:"siakng_cc" default:"2jNeTbVCFfkPIcnUkzwrVw" extensions:"x-order=01"`
	Mojavi   string `json:"mojavi" default:"UrdBjDansj/s95fYW58TfQ" extensions:"x-order=02"`
}

type LoginRequest struct {
	Username string `json:"username" default:"username" extensions:"x-order=01"`
	Password string `json:"password" default:"password" extensions:"x-order=02"`
}

// @Summary		login account
// @Description	get login cookie for other requests
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param loginRequest body LoginRequest true "login request"
// @Success		200	{object}	Response[CookieData]
// @Failure		401	{object}	Response[any]
// @Router		/login [post]
func Login(c *fiber.Ctx) error {
	var p LoginRequest

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
		return c.Status(fiber.StatusUnauthorized).JSON(Response[any]{
			Status:  401,
			Message: "Authentication failed",
			Data:    nil,
		})
	}

	resp, err = client.Get("https://academic.ui.ac.id/main/Authentication/ChangeRole")
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(Response[any]{
			Status:  401,
			Message: "Authentication failed",
			Data:    nil,
		})
	}

	for h := range resp.Header {
		if !strings.Contains(h, "Access-Control") {
			c.Set(h, resp.Header.Get(h))
		}
	}

	siakCookie := ""
	mojaviCookie := ""
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

		if cookie.Name == "siakng_cc" {
			siakCookie = cookie.Value
		}

		if cookie.Name == "Mojavi" {
			mojaviCookie = cookie.Value
		}

		c.Cookie(cookies)
	}

	return c.Status(resp.StatusCode).JSON(Response[CookieData]{
		Status:  200,
		Message: "Authentication success. Please use given cookie as X-MOJAVI and X-SIAKNG-CC for next requests",
		Data: CookieData{
			SiakNGCC: siakCookie,
			Mojavi:   mojaviCookie,
		},
	})
}

type Handler func(c *fiber.Ctx, response *http.Response) error

func BaseHandler(url string, next Handler) func(c *fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		req, err := MakeRequestor(c)
		if err != nil {
			return err
		}

		resp, err := req.Get(url)
		if err != nil {
			var e *SiakError
			if errors.As(err, &e) {
				return c.Status(e.Status).JSON(Response[any]{
					Status:  e.Status,
					Message: e.Message,
					Data:    nil,
				})
			}
			return err
		}

		return next(c, resp)
	}
}

// @Summary		Show an account
// @Description	get string by ID
// @Tags			accounts
// @Accept			json
// @Produce		json
// @Param			id	path		int	true	"Account ID"
// @Success		200	{object}	map[string]string
// @Router			/accounts/{id} [get]
func Home(c *fiber.Ctx, resp *http.Response) error {
	data, err := siaklib.ParseWelcomePage(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[siaklib.Homepage]{
		Status:  200,
		Message: "OK",
		Data:    *data,
	})
}

func Me(c *fiber.Ctx, resp *http.Response) error {
	data, err := siaklib.ParseNav(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[siaklib.UserInfo]{
		Status:  200,
		Message: "OK",
		Data:    *data,
	})
}

func AcademicSummary(c *fiber.Ctx, resp *http.Response) error {
	data, err := siaklib.ParseAcademicSummaryPage(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[siaklib.StudentSummary]{
		Status:  200,
		Message: "OK",
		Data:    *data,
	})
}

func AcademicHistory(c *fiber.Ctx, resp *http.Response) error {
	data, err := siaklib.ParseAcademicHistoryPage(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[[]siaklib.SemesterScore]{
		Status:  200,
		Message: "OK",
		Data:    *data,
	})
}

func UserPhoto(c *fiber.Ctx, resp *http.Response) error {
	for key, value := range resp.Header {
		if !strings.Contains(key, "Access-Control") {
			c.Set(key, strings.Join(value, ", "))
		}
	}
	return c.SendStream(resp.Body)
}

func CourseClasses(c *fiber.Ctx, resp *http.Response) error {
	data, err := siaklib.ParseCourseClasses(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[[]siaklib.Course]{
		Status:  200,
		Message: "OK",
		Data:    data,
	})
}

func CourseComponent(c *fiber.Ctx) error {
	courseId := c.Params("courseId")
	req, err := MakeRequestor(c)
	if err != nil {
		return err
	}

	resp, err := req.Get(fmt.Sprintf("https://academic.ui.ac.id/main/Academic/ScoreDetail?cc=%s", courseId))
	if err != nil {
		var e *SiakError
		if errors.As(err, &e) {
			return c.Status(e.Status).JSON(Response[any]{
				Status:  e.Status,
				Message: e.Message,
				Data:    nil,
			})
		}
		return err
	}

	data, err := siaklib.ParseCourseDetail(resp.Body)
	if err != nil {
		return err
	}

	return c.JSON(Response[[]siaklib.CourseComponent]{
		Status:  200,
		Message: "OK",
		Data:    data,
	})
}
