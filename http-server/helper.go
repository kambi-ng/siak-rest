package httpserver

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
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

const (
	sectigoCertFile = "sectigo.crt"
)

func MakeTlsTransport() (*http.Transport, error) {
	rootCAs, _ := x509.SystemCertPool()
	if rootCAs == nil {
		rootCAs = x509.NewCertPool()
	}

	// Read in the cert file
	certs, err := ioutil.ReadFile(sectigoCertFile)
	if err != nil {
		return nil, err
	}

	// Append our cert to the system pool
	if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
		return nil, errors.New("failed to add cert")
	}

	// Trust the augmented cert pool in our client
	config := &tls.Config{
		RootCAs: rootCAs,
	}
	tr := &http.Transport{TLSClientConfig: config}

	return tr, nil
}

func MakeRequestor(c *fiber.Ctx) (*http.Client, error) {
	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}

	tr, err := MakeTlsTransport()
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Transport: tr,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			if strings.Contains(req.URL.Path, "Authentication") {
				return &SiakError{Status: fiber.StatusUnauthorized, Message: "Session has expired"}
			}

			return nil
		},
		Jar: jar,
	}

	u, _ := url.Parse("https://academic.ui.ac.id")
	headers := c.GetReqHeaders()
	jar.SetCookies(u, []*http.Cookie{
		{
			Name:  "Mojavi",
			Value: headers["X-Mojavi"],
		},
		{
			Name:  "siakng_cc",
			Value: headers["X-Siakng-Cc"],
		},
	})

	return client, nil
}
