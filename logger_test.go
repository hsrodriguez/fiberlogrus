package fiberlogrus

import (
	"bytes"
	"log"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	s := ""
	buf := bytes.NewBufferString(s)

	logger := logrus.New()
	logger.SetOutput(buf)

	app.Use(New(
		Config{
			Logger: logger,
			Tags: []string{
				TagMethod,
				TagStatus,
				AttachKeyTag(TagLocals, "loc"),
				AttachKeyTag(TagResHeader, "custom-header"),
			},
		}))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Append("custom-header", "custom-header-value")
		c.Locals("loc", "val")
		return c.SendString("random string")
	})
	resp, err := app.Test(httptest.NewRequest(fiber.MethodGet, "/", nil))

	log.Println(buf.String())

	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)
	require.Contains(t, buf.String(), "method=GET")
	require.Contains(t, buf.String(), "status=200")
	require.Contains(t, buf.String(), "respHeader=custom-header-value")
	require.Contains(t, buf.String(), "locals=val")
}

func TestStringLogger(t *testing.T) {
	t.Parallel()
	app := fiber.New()

	s := ""
	buf := bytes.NewBufferString(s)

	loggerstd := logrus.New()
	loggerstd.Out = buf
	// loggerstd.Formatter = &logrus.JSONFormatter{PrettyPrint: true}

	app.Use(New(
		Config{
			Logger: loggerstd,
			Tags: []string{
				TagMethod,
				TagReqBody,
				TagReqBodyString,
				TagReqHeaders,
				TagReqHeadersString,
				TagResBody,
				TagResBodyString,
				TagResHeaders,
				TagResHeadersString,
			},
		},
	))

	app.Get("/", func(c *fiber.Ctx) error {
		c.Append("custom-response-header", "custom-response-header-value")
		return c.SendString("random string")
	})
	reader := strings.NewReader("number=2")
	req:=httptest.NewRequest(fiber.MethodGet, "/", reader)
	req.Header.Add("custom-request-header", "custom-request-header-value")
	resp, err := app.Test(req)

	log.Println(buf.String())

	require.Contains(t, buf.String(), "Custom-Request-Header=custom-request-header-value")
	require.Contains(t, buf.String(), "reqBodyString=\"number=2\"")
	require.Contains(t, buf.String(), "Custom-Response-Header=custom-response-header-value")
	require.Contains(t, buf.String(), "resBodyString=\"random string\"")

	require.NoError(t, err)
	require.Equal(t, fiber.StatusOK, resp.StatusCode)
}
