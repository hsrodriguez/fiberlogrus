package fiberlogrus

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

// FuncTag is a function used to populate logrus field
type FuncTag func(c *fiber.Ctx, d *data) any

// predefined FuncTag functions
var (
	FuncTagReferer FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Get(fiber.HeaderReferer)
	}
	FuncTagProtocol FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Protocol()
	}
	FuncTagPort FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Port()
	}
	FuncTagIP FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.IP()
	}
	FuncTagIPs FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Get(fiber.HeaderXForwardedFor)
	}
	FuncTagHost FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Hostname()
	}
	FuncTagPath FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Path()
	}
	FuncTagURL FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.OriginalURL()
	}
	FuncTagUA FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Get(fiber.HeaderUserAgent)
	}
	FuncTagReqBody FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Body()
	}
	FuncTagBytesReceived FuncTag = func(c *fiber.Ctx, d *data) any {
		return len(c.Request().Body())
	}
	FuncTagBytesSent FuncTag = func(c *fiber.Ctx, d *data) any {
		if c.Response().Header.ContentLength() < 0 {
			return 0
		}
		return len(c.Response().Body())
	}
	FuncTagRoute FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Route().Path
	}
	FuncTagResBody FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Response().Body()
	}
	FuncTagReqHeaders FuncTag = func(c *fiber.Ctx, d *data) any {
		reqHeaders := make([]string, 0)
		for k, v := range c.GetReqHeaders() {
			reqHeaders = append(reqHeaders, k+"="+v)
		}
		return []byte(strings.Join(reqHeaders, "&"))
	}
	FuncTagQueryStringParams FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Request().URI().QueryArgs().String()
	}
	FuncTagStatus FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Response().StatusCode()
	}
	FuncTagMethod FuncTag = func(c *fiber.Ctx, d *data) any {
		return c.Method()
	}
	FuncTagPid FuncTag = func(c *fiber.Ctx, d *data) any {
		return d.pid
	}
	FuncTagLatency FuncTag = func(c *fiber.Ctx, d *data) any {
		return d.end.Sub(d.start).String()
	}
	FuncTagResHeaders = func(c *fiber.Ctx, d *data) any {
		resHeaders := make([]string, 0)
		for k, v := range c.GetRespHeaders() {
			resHeaders = append(resHeaders, fmt.Sprintf("%s=%s", k, v))
		}
		return []byte(strings.Join(resHeaders, "&"))
	}
	FuncTagReqHeadersString = func(c *fiber.Ctx, d *data) any {
		reqHeaders := make([]string, 0)
		for k, v := range c.GetReqHeaders() {
			reqHeaders = append(reqHeaders, fmt.Sprintf("%s=%s", k, v))
		}
		return strings.Join(reqHeaders, "&")
	}
	FuncTagReqBodyString = func(c *fiber.Ctx, d *data) any {
		return string(c.Body())
	}
	FuncTagResHeadersString = func(c *fiber.Ctx, d *data) any {
		resHeaders := make([]string, 0)
		for k, v := range c.GetRespHeaders() {
			resHeaders = append(resHeaders, fmt.Sprintf("%s=%s", k, v))
		}
		return strings.Join(resHeaders, "&")
	}
	FuncTagResBodyString = func(c *fiber.Ctx, d *data) any {
		return string(c.Response().Body())
	}

	FuncTagReqHeader = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			return c.Get(extra)
		}
	}
	FuncTagResHeader = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			return c.GetRespHeader(extra)
		}
	}
	FuncTagQuery = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			return c.Query(extra)
		}
	}
	FuncTagForm = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			return c.FormValue(extra)
		}
	}
	FuncTagCookie = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			return c.Cookies(extra)
		}
	}
	FuncTagLocals = func(extra string) FuncTag {
		return func(c *fiber.Ctx, d *data) any {
			switch v := c.Locals(extra).(type) {
			case []byte:
				return string(v)
			case string:
				return v
			case nil:
				return nil
			default:
				return fmt.Sprintf("%v", v)
			}
		}
	}
)

// attached keyTag separator
const sep string = ":"

// getFuncTagMap selects functions to be used for logrus fields population
func getFuncTagMap(cfg Config, d *data) map[string]FuncTag {
	m := make(map[string]FuncTag, len(cfg.Tags))
	for _, t := range cfg.Tags {
		switch t {
		case TagReferer:
			m[TagReferer] = FuncTagReferer
		case TagProtocol:
			m[TagProtocol] = FuncTagProtocol
		case TagPort:
			m[TagPort] = FuncTagPort
		case TagIP:
			m[TagIP] = FuncTagIP
		case TagIPs:
			m[TagIPs] = FuncTagIPs
		case TagHost:
			m[TagHost] = FuncTagHost
		case TagPath:
			m[TagPath] = FuncTagPath
		case TagURL:
			m[TagURL] = FuncTagURL
		case TagUA:
			m[TagUA] = FuncTagUA
		case TagReqBody:
			m[TagReqBody] = FuncTagReqBody
		case TagBytesReceived:
			m[TagBytesReceived] = FuncTagBytesReceived
		case TagBytesSent:
			m[TagBytesSent] = FuncTagBytesSent
		case TagRoute:
			m[TagRoute] = FuncTagRoute
		case TagResBody:
			m[TagResBody] = FuncTagResBody
		case TagReqHeaders:
			m[TagReqHeaders] = FuncTagReqHeaders
		case TagQueryStringParams:
			m[TagQueryStringParams] = FuncTagQueryStringParams
		case TagStatus:
			m[TagStatus] = FuncTagStatus
		case TagMethod:
			m[TagMethod] = FuncTagMethod
		case TagPid:
			m[TagPid] = FuncTagPid
		case TagLatency:
			m[TagLatency] = FuncTagLatency
		case TagResHeaders:
			m[TagResHeaders] = FuncTagResHeaders
		case TagReqHeadersString:
			m[TagReqHeadersString] = FuncTagReqHeadersString
		case TagReqBodyString:
			m[TagReqBodyString] = FuncTagReqBodyString
		case TagResHeadersString:
			m[TagResHeadersString] = FuncTagResHeadersString
		case TagResBodyString:
			m[TagResBodyString] = FuncTagResBodyString
		default:
			for _, v := range KeyTags {
				if strings.Contains(t, v) {
					a := strings.Split(t, sep)
					switch a[0] {
					case TagReqHeader:
						m[TagReqHeader] = FuncTagReqHeader(a[1])
					case TagResHeader:
						m[TagResHeader] = FuncTagResHeader(a[1])
					case TagQuery:
						m[TagQuery] = FuncTagQuery(a[1])
					case TagForm:
						m[TagForm] = FuncTagForm(a[1])
					case TagCookie:
						m[TagCookie] = FuncTagCookie(a[1])
					case TagLocals:
						m[TagLocals] = FuncTagLocals(a[1])
					}

				}
			}
		}
	}
	return m
}
