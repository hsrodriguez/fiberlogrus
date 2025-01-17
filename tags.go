package fiberlogrus

import "fmt"

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

var CommonTags []string = []string{
	TagReferer,
	TagProtocol,
	TagPort,
	TagIP,
	TagIPs,
	TagHost,
	TagPath,
	TagURL,
	TagUA,
	TagReqBody,
	TagBytesReceived,
	TagBytesSent,
	TagRoute,
	TagResBody,
	TagReqHeaders,
	TagQueryStringParams,
	TagStatus,
	TagMethod,
	TagPid,
	TagLatency,
	TagResHeaders,
	TagReqHeadersString,
	TagReqBodyString,
	TagResHeadersString,
	TagResBodyString,
}

var KeyTags []string = []string{
	TagReqHeader,
	TagResHeader,
	TagQuery,
	TagForm,
	TagCookie,
	TagLocals,
}

// AttachKeyTag forms a string to access values stored as k-v pairs in format
//
// "keytag:key"
//
// Example:
//
// AttachKeyTag(TagLocals, "requestid")
//
// useing in Config.Tags will add a "locals:b367dbaf-7e1d-422c-97f0-ec4348c1bd0b" field in logger output
func AttachKeyTag(keyTag, key string) string {
	return fmt.Sprintf("%s:%s", keyTag, key)
}
