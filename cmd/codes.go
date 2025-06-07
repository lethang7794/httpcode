package cmd

import "fmt"

// HTTPCodeInfo contains detailed information about an HTTP status code
type HTTPCodeInfo struct {
	Description string
	Detail      string
	MDNLink     string
}

// HTTP status codes and their detailed information
var httpCodesInfo = map[int]HTTPCodeInfo{
	// 1xx Informational
	100: {
		Description: "Continue",
		Detail:      "The server has received the request headers and the client should proceed to send the request body.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/100",
	},
	101: {
		Description: "Switching Protocols",
		Detail:      "The requester has asked the server to switch protocols and the server has agreed to do so.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/101",
	},
	102: {
		Description: "Processing",
		Detail:      "The server has received and is processing the request, but no response is available yet.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/102",
	},
	103: {
		Description: "Early Hints",
		Detail:      "Used to return some response headers before final HTTP message.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/103",
	},

	// 2xx Success
	200: {
		Description: "OK",
		Detail:      "The request has succeeded. The information returned with the response depends on the method used in the request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/200",
	},
	201: {
		Description: "Created",
		Detail:      "The request has succeeded and a new resource has been created as a result. This is typically the response sent after POST requests, or some PUT requests.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/201",
	},
	202: {
		Description: "Accepted",
		Detail:      "The request has been received but not yet acted upon. It is noncommittal, since there is no way in HTTP to later send an asynchronous response indicating the outcome of the request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/202",
	},
	203: {
		Description: "Non-Authoritative Information",
		Detail:      "The returned metadata is not exactly the same as is available from the origin server, but is collected from a local or a third-party copy.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/203",
	},
	204: {
		Description: "No Content",
		Detail:      "The server successfully processed the request, but is not returning any content. Usually used as a response to a successful delete request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/204",
	},
	205: {
		Description: "Reset Content",
		Detail:      "The server successfully processed the request, asks that the requester reset its document view, and is not returning any content.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/205",
	},
	206: {
		Description: "Partial Content",
		Detail:      "The server is delivering only part of the resource due to a range header sent by the client. Used for resumable downloads and split downloads.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/206",
	},
	207: {
		Description: "Multi-Status",
		Detail:      "The message body that follows is by default an XML message and can contain a number of separate response codes, depending on how many sub-requests were made.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/207",
	},
	208: {
		Description: "Already Reported",
		Detail:      "The members of a DAV binding have already been enumerated in a preceding part of the (multistatus) response, and are not being included again.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/208",
	},
	226: {
		Description: "IM Used",
		Detail:      "The server has fulfilled a request for the resource, and the response is a representation of the result of one or more instance-manipulations applied to the current instance.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/226",
	},

	// 3xx Redirection
	300: {
		Description: "Multiple Choices",
		Detail:      "The request has more than one possible response. The user-agent or user should choose one of them.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/300",
	},
	301: {
		Description: "Moved Permanently",
		Detail:      "The URL of the requested resource has been changed permanently. The new URL is given in the response.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301",
	},
	302: {
		Description: "Found",
		Detail:      "The URI of requested resource has been changed temporarily. Further changes in the URI might be made in the future.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/302",
	},
	303: {
		Description: "See Other",
		Detail:      "The server sent this response to direct the client to get the requested resource at another URI with a GET request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/303",
	},
	304: {
		Description: "Not Modified",
		Detail:      "This is used for caching purposes. It tells the client that the response has not been modified, so the client can continue to use the same cached version of the response.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/304",
	},
	305: {
		Description: "Use Proxy",
		Detail:      "Defined in a previous version of the HTTP specification to indicate that a requested response must be accessed by a proxy. It has been deprecated due to security concerns regarding in-band configuration of a proxy.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/305",
	},
	307: {
		Description: "Temporary Redirect",
		Detail:      "The server sends this response to direct the client to get the requested resource at another URI with the same method that was used in the prior request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/307",
	},
	308: {
		Description: "Permanent Redirect",
		Detail:      "This means that the resource is now permanently located at another URI, specified by the Location: HTTP Response header.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/308",
	},

	// 4xx Client Errors
	400: {
		Description: "Bad Request",
		Detail:      "The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/400",
	},
	401: {
		Description: "Unauthorized",
		Detail:      "Although the HTTP standard specifies 'unauthorized', semantically this response means 'unauthenticated'. That is, the client must authenticate itself to get the requested response.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401",
	},
	402: {
		Description: "Payment Required",
		Detail:      "This response code is reserved for future use. The initial aim for creating this code was using it for digital payment systems, however this status code is used very rarely and no standard convention exists.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/402",
	},
	403: {
		Description: "Forbidden",
		Detail:      "The client does not have access rights to the content; that is, it is unauthorized, so the server is refusing to give the requested resource. Unlike 401, the client's identity is known to the server.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/403",
	},
	404: {
		Description: "Not Found",
		Detail:      "The server can not find the requested resource. In the browser, this means the URL is not recognized. In an API, this can also mean that the endpoint is valid but the resource itself does not exist.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/404",
	},
	405: {
		Description: "Method Not Allowed",
		Detail:      "The request method is known by the server but is not supported by the target resource. For example, an API may not allow DELETE a resource.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/405",
	},
	406: {
		Description: "Not Acceptable",
		Detail:      "This response is sent when the web server, after performing server-driven content negotiation, doesn't find any content that conforms to the criteria given by the user agent.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/406",
	},
	407: {
		Description: "Proxy Authentication Required",
		Detail:      "This is similar to 401 but authentication is needed to be done by a proxy.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/407",
	},
	408: {
		Description: "Request Timeout",
		Detail:      "This response is sent on an idle connection by some servers, even without any previous request by the client. It means that the server would like to shut down this unused connection.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/408",
	},
	409: {
		Description: "Conflict",
		Detail:      "This response is sent when a request conflicts with the current state of the server.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/409",
	},
	410: {
		Description: "Gone",
		Detail:      "This response is sent when the requested content has been permanently deleted from server, with no forwarding address. Clients are expected to remove their caches and links to the resource.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/410",
	},
	411: {
		Description: "Length Required",
		Detail:      "Server rejected the request because the Content-Length header field is not defined and the server requires it.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/411",
	},
	412: {
		Description: "Precondition Failed",
		Detail:      "The client has indicated preconditions in its headers which the server does not meet.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/412",
	},
	413: {
		Description: "Payload Too Large",
		Detail:      "Request entity is larger than limits defined by server; the server might close the connection or return an Retry-After header field.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/413",
	},
	414: {
		Description: "URI Too Long",
		Detail:      "The URI requested by the client is longer than the server is willing to interpret.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/414",
	},
	415: {
		Description: "Unsupported Media Type",
		Detail:      "The media format of the requested data is not supported by the server, so the server is rejecting the request.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/415",
	},
	416: {
		Description: "Range Not Satisfiable",
		Detail:      "The range specified by the Range header field in the request can't be fulfilled; it's possible that the range is outside the size of the target URI's data.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/416",
	},
	417: {
		Description: "Expectation Failed",
		Detail:      "This response code means the expectation indicated by the Expect request header field can't be met by the server.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/417",
	},
	418: {
		Description: "I'm a teapot",
		Detail:      "The server refuses the attempt to brew coffee with a teapot. This code was defined as an April Fools' joke in 1998.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/418",
	},
	421: {
		Description: "Misdirected Request",
		Detail:      "The request was directed at a server that is not able to produce a response. This can be sent by a server that is not configured to produce responses for the combination of scheme and authority that are included in the request URI.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/421",
	},
	422: {
		Description: "Unprocessable Entity",
		Detail:      "The request was well-formed but was unable to be followed due to semantic errors. Commonly used with validation errors in APIs.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/422",
	},
	423: {
		Description: "Locked",
		Detail:      "The resource that is being accessed is locked. Used in WebDAV.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/423",
	},
	424: {
		Description: "Failed Dependency",
		Detail:      "The request failed due to failure of a previous request. Used in WebDAV.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/424",
	},
	425: {
		Description: "Too Early",
		Detail:      "Indicates that the server is unwilling to risk processing a request that might be replayed.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/425",
	},
	426: {
		Description: "Upgrade Required",
		Detail:      "The server refuses to perform the request using the current protocol but might be willing to do so after the client upgrades to a different protocol.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/426",
	},
	428: {
		Description: "Precondition Required",
		Detail:      "The origin server requires the request to be conditional. This response is intended to prevent the 'lost update' problem, where a client GETs a resource's state, modifies it, and PUTs it back to the server, when meanwhile a third party has modified the state on the server, leading to a conflict.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/428",
	},
	429: {
		Description: "Too Many Requests",
		Detail:      "The user has sent too many requests in a given amount of time ('rate limiting'). Often used for API rate limiting.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/429",
	},
	431: {
		Description: "Request Header Fields Too Large",
		Detail:      "The server is unwilling to process the request because its header fields are too large. The request may be resubmitted after reducing the size of the request header fields.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/431",
	},
	451: {
		Description: "Unavailable For Legal Reasons",
		Detail:      "The user-agent requested a resource that cannot legally be provided, such as a web page censored by a government.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/451",
	},

	// 5xx Server Errors
	500: {
		Description: "Internal Server Error",
		Detail:      "The server has encountered a situation it doesn't know how to handle. A generic error message, given when an unexpected condition was encountered and no more specific message is suitable.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/500",
	},
	501: {
		Description: "Not Implemented",
		Detail:      "The request method is not supported by the server and cannot be handled. The only methods that servers are required to support (and therefore that must not return this code) are GET and HEAD.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/501",
	},
	502: {
		Description: "Bad Gateway",
		Detail:      "This error response means that the server, while working as a gateway to get a response needed to handle the request, got an invalid response.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/502",
	},
	503: {
		Description: "Service Unavailable",
		Detail:      "The server is not ready to handle the request. Common causes are a server that is down for maintenance or that is overloaded. Note that together with this response, a user-friendly page explaining the problem should be sent.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/503",
	},
	504: {
		Description: "Gateway Timeout",
		Detail:      "This error response is given when the server is acting as a gateway and cannot get a response in time.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/504",
	},
	505: {
		Description: "HTTP Version Not Supported",
		Detail:      "The HTTP version used in the request is not supported by the server.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/505",
	},
	506: {
		Description: "Variant Also Negotiates",
		Detail:      "The server has an internal configuration error: the chosen variant resource is configured to engage in transparent content negotiation itself, and is therefore not a proper end point in the negotiation process.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/506",
	},
	507: {
		Description: "Insufficient Storage",
		Detail:      "The server is unable to store the representation needed to complete the request. Used in WebDAV.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/507",
	},
	508: {
		Description: "Loop Detected",
		Detail:      "The server detected an infinite loop while processing the request. Used in WebDAV.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/508",
	},
	510: {
		Description: "Not Extended",
		Detail:      "Further extensions to the request are required for the server to fulfill it.",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/510",
	},
	511: {
		Description: "Network Authentication Required",
		Detail:      "The client needs to authenticate to gain network access. Intended for use by intercepting proxies used to control access to the network (e.g., 'captive portals' used to require agreement to Terms of Service before granting full Internet access via a Wi-Fi hotspot).",
		MDNLink:     "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/511",
	},
}

// Helper function to get just the HTTP codes map
var httpCodes = func() map[int]string {
	codes := make(map[int]string)
	for code, info := range httpCodesInfo {
		codes[code] = info.Description
	}
	return codes
}()

// Helper function to look up a specific HTTP status code
func lookupCode(code int) {
	if info, exists := httpCodesInfo[code]; exists {
		fmt.Printf("%d: %s\n", code, info.Description)
		fmt.Printf("\nDetail: %s\n", info.Detail)
		fmt.Printf("\nMDN Documentation: %s\n", info.MDNLink)
	} else {
		fmt.Printf("HTTP status code %d not found\n", code)
	}
}
