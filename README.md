# Logrus logger middleware [![Mentioned in Awesome Fiber](https://awesome.re/mentioned-badge-flat.svg)](https://github.com/gofiber/awesome-fiber)
Logger middleware for [Fiber](https://github.com/gofiber/fiber) that logs HTTP request/response details.


Use your configured `logrus` logger instance or global logrus instance to handle logging in a structured way.

## Table of Contents
- [Getting started](#getting-started)
- [Signatures](#signatures)
- [Examples](#examples)

## Getting started
```bash
$ go get github.com/hsrodriguez/fiberlogrus
```
## Signatures
```go
func New(config ...Config) fiber.Handler
```
## Examples
Import required packages
```go
import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/hsrodriguez/fiberlogrus"
)
```
### Default config
Using with a default config, it will call global logrus instance to log the requests
```go
app := fiber.New()

app.Use(fiberlogrus.New())
```
```go
// ConfigDefault is the default config
var ConfigDefault Config = Config{
	Logger: nil,
	Tags: []string{
		TagStatus,
		TagLatency,
		TagMethod,
		TagPath,
	},
}
```
### Use logger instance and configure tags
```go
logger := logrus.New()
// you can also provide logger with a desired formatter
// logger.SetFormatter(&logrus.JSONFormatter{})

app.Use(
	fiberlogrus.New(
		fiberlogrus.Config{
			Logger: logger,
			Tags: []string{
				// add method field
				fiberlogrus.TagMethod,
				// add status field
				fiberlogrus.TagStatus,
				// add value from locals
				AttachKeyTag(TagLocals, "requestid"),
				// add certain header
				AttachKeyTag(TagReqHeader, "custom-header"),
			},
		},
	),
)
```
### All supported common tags example
```go
package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"
	"github.com/mikhail-bigun/fiberlogrus"
)

func main() {
	app := fiber.New()

	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	app.Use(
		fiberlogrus.New(
			fiberlogrus.Config{
				Logger: logger,
				Tags: fiberlogrus.CommonTags,
			}))
	
	app.Get("/", func(c *fiber.Ctx) error { return c.SendStatus(fiber.StatusOK) })
	logger.Fatal(f.Listen(":8080"))
}
```

### Supported tags
#### Common
```go
// Common Tags
const (
	// request referer
	TagReferer = "referer"
	// request protocol
	TagProtocol = "protocol"
	// request port
	TagPort = "port"
	// request ip
	TagIP = "ip"
	// request ips
	TagIPs = "ips"
	// request host
	TagHost = "host"
	// request path
	TagPath = "path"
	// request url
	TagURL = "url"
	// request user-agent
	TagUA = "ua"
	// request body
	TagReqBody = "reqBody"
	// request body bytes length
	TagBytesReceived = "bytesReceived"
	// response bytes length
	TagBytesSent = "bytesSent"
	// request route
	TagRoute = "route"
	// response body
	TagResBody = "resBody"
	// request headers
	TagReqHeaders = "reqHeaders"
	// request query parameters
	TagQueryStringParams = "queryParams"
	// response status
	TagStatus = "status"
	// request method
	TagMethod = "method"
	// fiber process id
	TagPid = "pid"
	// request latency
	TagLatency = "latency"
	// response headers
	TagResHeaders = "resHeaders"
	// request headers string
	TagReqHeadersString = "reqHeadersString"
	// request body string
	TagReqBodyString = "reqBodyString"
	// response headers string
	TagResHeadersString = "resHeadersString"
	// response body string
	TagResBodyString = "resBodyString"
)
```
#### Key
```go
// Key Tags
const (
	// request specified header
	TagReqHeader = "reqHeader"
	// response specified header
	TagResHeader = "respHeader"
	// request specified query
	TagQuery = "query"
	// request specified form value
	TagForm = "form"
	// request specified cookie value
	TagCookie = "cookie"
	// request specified locals value
	TagLocals = "locals"
)
```