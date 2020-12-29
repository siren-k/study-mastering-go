## _12장에서 다룰 주제_

---
> * TCP/IP란 무엇이고 이것이 중요한 이유
> * IPv4와 IPv6 프로토콜
> * Netcat 커맨드라인 유틸리티
> * net/http 패키지
> * http.Response, http.Request, http.Transport 구조체
> * Go 언어로 웹 서버 작성하기
> * Go 언어로 웹 클라이언트 작성하기
> * Go 언어로 웹사이트 만들기
> * Wireshark와 tshark
> * 서버나 클라이언트에서 응답하는데 너무 오래 걸리는 HTTP 연결 끊기

---
> ## _net/http, net, http.RoundTripper_
>> net/http 패키지는 강력한 웹 서버와 클라이언트를 개발하는데 필요한 기능을 제공한다. http.Get() 메소드는
>> 각각 HTTP와 HTTPS 용청을 보내는데 사용되고, http.ListenAndServe() 함수에 서비스를 제공할 IP 주소와
>> TCP 포트와 들어온 입력을 처리할 함수들을 지정하면 웹 서버를 생성할 수 있다.
>>
>> http.RoundTripper는 일종의 인터페이스로서 Go 프로그램이 HTTP 트랜잭션을 수행할 수 있도록 보장해주는데
>> 유용하다. 쉽게 말해 주어진 http.Request에 대해 http.Response를 받게 만들 수 있다.
>>
>>> ### _http.Response 타입_
>>> http.Response 구조체는 https://golang.org/src/net/http/response.go 파일에 다음과 같이
>>> 정의되어 있다. http.Response 타입의 목적은 HTTP 요청에 대한 응답을 표현하는 것이다. 소스 파일을
>>> 보면 각각의 필드에 대한 설명이 자세히 나와 있는데 표준 Go 라이브러리에서 정의하는 struct 타입은 대부분
>>> 이렇게 주석이 자세히 달려있다.
>>> ---
>>> ```
>>> // Response represents the response from an HTTP request.
>>> //
>>> // The Client and Transport return Responses from servers once
>>> // the response headers have been received. The response body
>>> // is streamed on demand as the Body field is read.
>>> type Response struct {
>>>     Status     string // e.g. "200 OK"
>>>     StatusCode int    // e.g. 200
>>>     Proto      string // e.g. "HTTP/1.0"
>>>     ProtoMajor int    // e.g. 1
>>>     ProtoMinor int    // e.g. 0
>>>
>>>     // Header maps header keys to values. If the response had multiple
>>>     // headers with the same key, they may be concatenated, with comma
>>>     // delimiters.  (RFC 7230, section 3.2.2 requires that multiple headers
>>>     // be semantically equivalent to a comma-delimited sequence.) When
>>>     // Header values are duplicated by other fields in this struct (e.g.,
>>>     // ContentLength, TransferEncoding, Trailer), the field values are
>>>     // authoritative.
>>>     //
>>>     // Keys in the map are canonicalized (see CanonicalHeaderKey).
>>>     Header Header
>>>     
>>>     // Body represents the response body.
>>>     //
>>>     // The response body is streamed on demand as the Body field
>>>     // is read. If the network connection fails or the server
>>>     // terminates the response, Body.Read calls return an error.
>>>     //
>>>	    // The http Client and Transport guarantee that Body is always
>>>     // non-nil, even on responses without a body or responses with
>>>	    // a zero-length body. It is the caller's responsibility to
>>>     // close Body. The default HTTP client's Transport may not
>>>     // reuse HTTP/1.x "keep-alive" TCP connections if the Body is
>>>     // not read to completion and closed.
>>>     //
>>>     // The Body is automatically dechunked if the server replied
>>>     // with a "chunked" Transfer-Encoding.
>>>     //
>>>     // As of Go 1.12, the Body will also implement io.Writer
>>>     // on a successful "101 Switching Protocols" response,
>>>     // as used by WebSockets and HTTP/2's "h2c" mode.
>>>     Body io.ReadCloser
>>>
>>>     // ContentLength records the length of the associated content. The
>>>     // value -1 indicates that the length is unknown. Unless Request.Method
>>>	    // is "HEAD", values >= 0 indicate that the given number of bytes may
>>>     // be read from Body.
>>>     ContentLength int64
>>>
>>>     // Contains transfer encodings from outer-most to inner-most. Value is
>>>     // nil, means that "identity" encoding is used.
>>>     TransferEncoding []string
>>>
>>>     // Close records whether the header directed that the connection be
>>>     // closed after reading Body. The value is advice for clients: neither
>>>     // ReadResponse nor Response.Write ever closes a connection.
>>>     Close bool
>>>
>>>     // Uncompressed reports whether the response was sent compressed but
>>>     // was decompressed by the http package. When true, reading from
>>>     // Body yields the uncompressed content instead of the compressed
>>>     // content actually set from the server, ContentLength is set to -1,
>>>     // and the "Content-Length" and "Content-Encoding" fields are deleted
>>>     // from the responseHeader. To get the original response from
>>>     // the server, set Transport.DisableCompression to true.
>>>     Uncompressed bool
>>>
>>>     // Trailer maps trailer keys to values in the same
>>>     // format as Header.
>>>     //
>>>     // The Trailer initially contains only nil values, one for
>>>     // each key specified in the server's "Trailer" header
>>>     // value. Those values are not added to Header.
>>>     //
>>>     // Trailer must not be accessed concurrently with Read calls
>>>     // on the Body.
>>>     //
>>>     // After Body.Read has returned io.EOF, Trailer will contain
>>>     // any trailer values sent by the server.
>>>     Trailer Header
>>>
>>>     // Request is the request that was sent to obtain this Response.
>>>     // Request's Body is nil (having already been consumed).
>>>     // This is only populated for Client requests.
>>>     Request *Request
>>>
>>>     // TLS contains information about the TLS connection on which the
>>>     // response was received. It is nil for unencrypted responses.
>>>     // The pointer is shared between responses and should not be
>>>     // modified.
>>>     TLS *tls.ConnectionState
>>> }```
>>
>>> ### _http.Request 타입_
>>> http.Request 타입의 목적은 서버에서 받거나 클라이언트에서 보낸 HTTP 요청을 표현하는 것이다. http.Request
>>> 구조체는 https://golang.org/src/net/http/request.go 파일에 다음과 같이 정의돼 있다.
>>> ---
>>> ```
>>> // A Request represents an HTTP request received by a server
>>> // or to be sent by a client.
>>> //
>>> // The field semantics differ slightly between client and server
>>> // usage. In addition to the notes on the fields below, see the
>>> // documentation for Request.Write and RoundTripper.
>>> type Request struct {
>>>     // Method specifies the HTTP method (GET, POST, PUT, etc.).
>>>     // For client requests, an empty string means GET.
>>>     //
>>>     // Go's HTTP client does not support sending a request with
>>>     // the CONNECT method. See the documentation on Transport for
>>>     // details.
>>>     Method string

>>>     // URL specifies either the URI being requested (for server
>>>     // requests) or the URL to access (for client requests).
>>>     //
>>>     // For server requests, the URL is parsed from the URI
>>>     // supplied on the Request-Line as stored in RequestURI.  For
>>>     // most requests, fields other than Path and RawQuery will be
>>>     // empty. (See RFC 7230, Section 5.3)
>>>     //
>>>     // For client requests, the URL's Host specifies the server to
>>>     // connect to, while the Request's Host field optionally
>>>     // specifies the Host header value to send in the HTTP
>>>     // request.
>>>     URL *url.URL
>>>
>>>     // The protocol version for incoming server requests.
>>>     //
>>>     // For client requests, these fields are ignored. The HTTP
>>>     // client code always uses either HTTP/1.1 or HTTP/2.
>>>     // See the docs on Transport for details.
>>>     Proto      string // "HTTP/1.0"
>>>     ProtoMajor int    // 1
>>>     ProtoMinor int    // 0
>>>
>>>     // Header contains the request header fields either received
>>>     // by the server or to be sent by the client.
>>>     //
>>>     // If a server received a request with header lines,
>>>     //
>>>     //	Host: example.com
>>>     //	accept-encoding: gzip, deflate
>>>     //	Accept-Language: en-us
>>>     //	fOO: Bar
>>>     //	foo: two
>>>     //
>>>     // then
>>>     //
>>>     //	Header = map[string][]string{
>>>     //		"Accept-Encoding": {"gzip, deflate"},
>>>     //		"Accept-Language": {"en-us"},
>>>     //		"Foo": {"Bar", "two"},
>>>     //	}
>>>     //
>>>     // For incoming requests, the Host header is promoted to the
>>>     // Request.Host field and removed from the Header map.
>>>     //
>>>     // HTTP defines that header names are case-insensitive. The
>>>     // request parser implements this by using CanonicalHeaderKey,
>>>     // making the first character and any characters following a
>>>     // hyphen uppercase and the rest lowercase.
>>>     //
>>>     // For client requests, certain headers such as Content-Length
>>>     // and Connection are automatically written when needed and
>>>     // values in Header may be ignored. See the documentation
>>>     // for the Request.Write method.
>>>     Header Header
>>>     
>>>     // Body is the request's body.
>>>     //
>>>     // For client requests, a nil body means the request has no
>>>     // body, such as a GET request. The HTTP Client's Transport
>>>     // is responsible for calling the Close method.
>>>     //
>>>     // For server requests, the Request Body is always non-nil
>>>     // but will return EOF immediately when no body is present.
>>>     // The Server will close the request body. The ServeHTTP
>>>     // Handler does not need to.
>>>     Body io.ReadCloser
>>>
>>>     // GetBody defines an optional func to return a new copy of
>>>     // Body. It is used for client requests when a redirect requires
>>>     // reading the body more than once. Use of GetBody still
>>>     // requires setting Body.
>>>     //
>>>     // For server requests, it is unused.
>>>     GetBody func() (io.ReadCloser, error)
>>>
>>>     // ContentLength records the length of the associated content.
>>>     // The value -1 indicates that the length is unknown.
>>>     // Values >= 0 indicate that the given number of bytes may
>>>     // be read from Body.
>>>     //
>>>     // For client requests, a value of 0 with a non-nil Body is
>>>     // also treated as unknown.
>>>     ContentLength int64
>>>
>>>     // TransferEncoding lists the transfer encodings from outermost to
>>>     // innermost. An empty list denotes the "identity" encoding.
>>>     // TransferEncoding can usually be ignored; chunked encoding is
>>>     // automatically added and removed as necessary when sending and
>>>     // receiving requests.
>>>     TransferEncoding []string
>>>
>>>     // Close indicates whether to close the connection after
>>>     // replying to this request (for servers) or after sending this
>>>     // request and reading its response (for clients).
>>>     //
>>>     // For server requests, the HTTP server handles this automatically
>>>     // and this field is not needed by Handlers.
>>>     //
>>>     // For client requests, setting this field prevents re-use of
>>>     // TCP connections between requests to the same hosts, as if
>>>     // Transport.DisableKeepAlives were set.
>>>     Close bool
>>>
>>>     // For server requests, Host specifies the host on which the
>>>     // URL is sought. For HTTP/1 (per RFC 7230, section 5.4), this
>>>     // is either the value of the "Host" header or the host name
>>>     // given in the URL itself. For HTTP/2, it is the value of the
>>>     // ":authority" pseudo-header field.
>>>     // It may be of the form "host:port". For international domain
>>>     // names, Host may be in Punycode or Unicode form. Use
>>>     // golang.org/x/net/idna to convert it to either format if
>>>     // needed.
>>>     // To prevent DNS rebinding attacks, server Handlers should
>>>     // validate that the Host header has a value for which the
>>>     // Handler considers itself authoritative. The included
>>>     // ServeMux supports patterns registered to particular host
>>>     // names and thus protects its registered Handlers.
>>>     //
>>>     // For client requests, Host optionally overrides the Host
>>>     // header to send. If empty, the Request.Write method uses
>>>     // the value of URL.Host. Host may contain an international
>>>     // domain name.
>>>     Host string
>>>
>>>     // Form contains the parsed form data, including both the URL
>>>     // field's query parameters and the PATCH, POST, or PUT form data.
>>>     // This field is only available after ParseForm is called.
>>>     // The HTTP client ignores Form and uses Body instead.
>>>     Form url.Values
>>>
>>>     // PostForm contains the parsed form data from PATCH, POST
>>>     // or PUT body parameters.
>>>     //
>>>     // This field is only available after ParseForm is called.
>>>     // The HTTP client ignores PostForm and uses Body instead.
>>>     PostForm url.Values
>>>     
>>>     // MultipartForm is the parsed multipart form, including file uploads.
>>>     // This field is only available after ParseMultipartForm is called.
>>>     // The HTTP client ignores MultipartForm and uses Body instead.
>>>     MultipartForm *multipart.Form
>>>
>>>     // Trailer specifies additional headers that are sent after the request
>>>     // body.
>>>     //
>>>     // For server requests, the Trailer map initially contains only the
>>>     // trailer keys, with nil values. (The client declares which trailers it
>>>     // will later send.)  While the handler is reading from Body, it must
>>>     // not reference Trailer. After reading from Body returns EOF, Trailer
>>>     // can be read again and will contain non-nil values, if they were sent
>>>     // by the client.
>>>     //
>>>     // For client requests, Trailer must be initialized to a map containing
>>>     // the trailer keys to later send. The values may be nil or their final
>>>     // values. The ContentLength must be 0 or -1, to send a chunked request.
>>>     // After the HTTP request is sent the map values can be updated while
>>>     // the request body is read. Once the body returns EOF, the caller must
>>>     // not mutate Trailer.
>>>     //
>>>     // Few HTTP clients, servers, or proxies support HTTP trailers.
>>>     Trailer Header
>>>
>>>     // RemoteAddr allows HTTP servers and other software to record
>>>     // the network address that sent the request, usually for
>>>     // logging. This field is not filled in by ReadRequest and
>>>     // has no defined format. The HTTP server in this package
>>>     // sets RemoteAddr to an "IP:port" address before invoking a
>>>     // handler.
>>>     // This field is ignored by the HTTP client.
>>>     RemoteAddr string

>>>     // RequestURI is the unmodified request-target of the
>>>     // Request-Line (RFC 7230, Section 3.1.1) as sent by the client
>>>     // to a server. Usually the URL field should be used instead.
>>>     // It is an error to set this field in an HTTP client request.
>>>     RequestURI string
>>>
>>>     // TLS allows HTTP servers and other software to record
>>>     // information about the TLS connection on which the request
>>>     // was received. This field is not filled in by ReadRequest.
>>>     // The HTTP server in this package sets the field for
>>>     // TLS-enabled connections before invoking a handler;
>>>     // otherwise it leaves the field nil.
>>>     // This field is ignored by the HTTP client.
>>>     TLS *tls.ConnectionState
>>>
>>>     // Cancel is an optional channel whose closure indicates that the client
>>>     // request should be regarded as canceled. Not all implementations of
>>>     // RoundTripper may support Cancel.
>>>     //
>>>     // For server requests, this field is not applicable.
>>>     //
>>>     // Deprecated: Set the Request's context with NewRequestWithContext
>>>     // instead. If a Request's Cancel field and context are both
>>>     // set, it is undefined whether Cancel is respected.
>>>     Cancel <-chan struct{}
>>>
>>>     // Response is the redirect response which caused this request
>>>     // to be created. This field is only populated during client
>>>     // redirects.
>>>     Response *Response
>>>
>>>     // ctx is either the client or server context. It should only
>>>     // be modified via copying the whole Request using WithContext.
>>>     // It is unexported to prevent people from using Context wrong
>>>     // and mutating the contexts held by callers of the same request.
>>>     ctx context.Context
>>> }```
>>
>>> ### _http.Transport 타입_
>>> http.Transport 구조체는 https://golang.org/src/net/http/transport.go 파일에 다음과
>>> 같이 정의되어 있다. http.Transport는 상당히 방대한 필드로 구성된 복잡한 구조체라는 것을 알 수 있다.
>>> 그나마 다행인 것은 코드에서 http.Transport 구조체를 항상 사용하는 것이 아닌고 사용하더라도 모든 필드를
>>> 다루지 않아도 된다는 것이다.
>>> ---
>>> ```
>>> // Transport is an implementation of RoundTripper that supports HTTP,
>>> // HTTPS, and HTTP proxies (for either HTTP or HTTPS with CONNECT).
>>> //
>>> // By default, Transport caches connections for future re-use.
>>> // This may leave many open connections when accessing many hosts.
>>> // This behavior can be managed using Transport's CloseIdleConnections method
>>> // and the MaxIdleConnsPerHost and DisableKeepAlives fields.
>>> //
>>> // Transports should be reused instead of created as needed.
>>> // Transports are safe for concurrent use by multiple goroutines.
>>> //
>>> // A Transport is a low-level primitive for making HTTP and HTTPS requests.
>>> // For high-level functionality, such as cookies and redirects, see Client.
>>> //
>>> // Transport uses HTTP/1.1 for HTTP URLs and either HTTP/1.1 or HTTP/2
>>> // for HTTPS URLs, depending on whether the server supports HTTP/2,
>>> // and how the Transport is configured. The DefaultTransport supports HTTP/2.
>>> // To explicitly enable HTTP/2 on a transport, use golang.org/x/net/http2
>>> // and call ConfigureTransport. See the package docs for more about HTTP/2.
>>> //
>>> // Responses with status codes in the 1xx range are either handled
>>> // automatically (100 expect-continue) or ignored. The one
>>> // exception is HTTP status code 101 (Switching Protocols), which is
>>> // considered a terminal status and returned by RoundTrip. To see the
>>> // ignored 1xx responses, use the httptrace trace package's
>>> // ClientTrace.Got1xxResponse.
>>> //
>>> // Transport only retries a request upon encountering a network error
>>> // if the request is idempotent and either has no body or has its
>>> // Request.GetBody defined. HTTP requests are considered idempotent if
>>> // they have HTTP methods GET, HEAD, OPTIONS, or TRACE; or if their
>>> // Header map contains an "Idempotency-Key" or "X-Idempotency-Key"
>>> // entry. If the idempotency key value is a zero-length slice, the
>>> // request is treated as idempotent but the header is not sent on the
>>> // wire.
>>> type Transport struct {
>>>     idleMu       sync.Mutex
>>>     closeIdle    bool                                // user has requested to close all idle conns
>>>     idleConn     map[connectMethodKey][]*persistConn // most recently used at end
>>>     idleConnWait map[connectMethodKey]wantConnQueue  // waiting getConns
>>>     idleLRU      connLRU
>>>     
>>>     reqMu       sync.Mutex
>>>     reqCanceler map[cancelKey]func(error)
>>>     
>>>     altMu    sync.Mutex   // guards changing altProto only
>>>     altProto atomic.Value // of nil or map[string]RoundTripper, key is URI scheme
>>>     
>>>     connsPerHostMu   sync.Mutex
>>>     connsPerHost     map[connectMethodKey]int
>>>     connsPerHostWait map[connectMethodKey]wantConnQueue // waiting getConns
>>>     
>>>     // Proxy specifies a function to return a proxy for a given
>>>     // Request. If the function returns a non-nil error, the
>>>     // request is aborted with the provided error.
>>>     //
>>>     // The proxy type is determined by the URL scheme. "http",
>>>     // "https", and "socks5" are supported. If the scheme is empty,
>>>     // "http" is assumed.
>>>     //
>>>     // If Proxy is nil or returns a nil *URL, no proxy is used.
>>>     Proxy func(*Request) (*url.URL, error)
>>>     
>>>     // DialContext specifies the dial function for creating unencrypted TCP connections.
>>>     // If DialContext is nil (and the deprecated Dial below is also nil),
>>>     // then the transport dials using package net.
>>>     //
>>>     // DialContext runs concurrently with calls to RoundTrip.
>>>     // A RoundTrip call that initiates a dial may end up using
>>>     // a connection dialed previously when the earlier connection
>>>     // becomes idle before the later DialContext completes.
>>>     DialContext func(ctx context.Context, network, addr string) (net.Conn, error)
>>>
>>>     // Dial specifies the dial function for creating unencrypted TCP connections.
>>>     //
>>>     // Dial runs concurrently with calls to RoundTrip.
>>>     // A RoundTrip call that initiates a dial may end up using
>>>     // a connection dialed previously when the earlier connection
>>>     // becomes idle before the later Dial completes.
>>>     //
>>>     // Deprecated: Use DialContext instead, which allows the transport
>>>     // to cancel dials as soon as they are no longer needed.
>>>     // If both are set, DialContext takes priority.
>>>     Dial func(network, addr string) (net.Conn, error)
>>>
>>>     // DialTLSContext specifies an optional dial function for creating
>>>     // TLS connections for non-proxied HTTPS requests.
>>>     //
>>>     // If DialTLSContext is nil (and the deprecated DialTLS below is also nil),
>>>     // DialContext and TLSClientConfig are used.
>>>     //
>>>     // If DialTLSContext is set, the Dial and DialContext hooks are not used for HTTPS
>>>     // requests and the TLSClientConfig and TLSHandshakeTimeout
>>>     // are ignored. The returned net.Conn is assumed to already be
>>>     // past the TLS handshake.
>>>     DialTLSContext func(ctx context.Context, network, addr string) (net.Conn, error)
>>>
>>>     // DialTLS specifies an optional dial function for creating
>>>     // TLS connections for non-proxied HTTPS requests.
>>>     //
>>>     // Deprecated: Use DialTLSContext instead, which allows the transport
>>>     // to cancel dials as soon as they are no longer needed.
>>>     // If both are set, DialTLSContext takes priority.
>>>     DialTLS func(network, addr string) (net.Conn, error)
>>>
>>>     // TLSClientConfig specifies the TLS configuration to use with
>>>     // tls.Client.
>>>     // If nil, the default configuration is used.
>>>     // If non-nil, HTTP/2 support may not be enabled by default.
>>>     TLSClientConfig *tls.Config
>>>
>>>     // TLSHandshakeTimeout specifies the maximum amount of time waiting to
>>>     // wait for a TLS handshake. Zero means no timeout.
>>>     TLSHandshakeTimeout time.Duration
>>>
>>>     // DisableKeepAlives, if true, disables HTTP keep-alives and
>>>     // will only use the connection to the server for a single
>>>     // HTTP request.
>>>     //
>>>     // This is unrelated to the similarly named TCP keep-alives.
>>>     DisableKeepAlives bool
>>>
>>>     // DisableCompression, if true, prevents the Transport from
>>>     // requesting compression with an "Accept-Encoding: gzip"
>>>     // request header when the Request contains no existing
>>>     // Accept-Encoding value. If the Transport requests gzip on
>>>     // its own and gets a gzipped response, it's transparently
>>>     // decoded in the Response.Body. However, if the user
>>>     // explicitly requested gzip it is not automatically
>>>     // uncompressed.
>>>     DisableCompression bool
>>>
>>>     // MaxIdleConns controls the maximum number of idle (keep-alive)
>>>     // connections across all hosts. Zero means no limit.
>>>     MaxIdleConns int
>>>
>>>     // MaxIdleConnsPerHost, if non-zero, controls the maximum idle
>>>     // (keep-alive) connections to keep per-host. If zero,
>>>     // DefaultMaxIdleConnsPerHost is used.
>>>     MaxIdleConnsPerHost int
>>>
>>>     // MaxConnsPerHost optionally limits the total number of
>>>     // connections per host, including connections in the dialing,
>>>     // active, and idle states. On limit violation, dials will block.
>>>     //
>>>     // Zero means no limit.
>>>     MaxConnsPerHost int
>>>
>>>     // IdleConnTimeout is the maximum amount of time an idle
>>>     // (keep-alive) connection will remain idle before closing
>>>     // itself.
>>>     // Zero means no limit.
>>>     IdleConnTimeout time.Duration
>>>
>>>     // ResponseHeaderTimeout, if non-zero, specifies the amount of
>>>     // time to wait for a server's response headers after fully
>>>     // writing the request (including its body, if any). This
>>>     // time does not include the time to read the response body.
>>>     ResponseHeaderTimeout time.Duration
>>>
>>>     // ExpectContinueTimeout, if non-zero, specifies the amount of
>>>     // time to wait for a server's first response headers after fully
>>>     // writing the request headers if the request has an
>>>     // "Expect: 100-continue" header. Zero means no timeout and
>>>     // causes the body to be sent immediately, without
>>>     // waiting for the server to approve.
>>>     // This time does not include the time to send the request header.
>>>     ExpectContinueTimeout time.Duration
>>>
>>>     // TLSNextProto specifies how the Transport switches to an
>>>     // alternate protocol (such as HTTP/2) after a TLS ALPN
>>>     // protocol negotiation. If Transport dials an TLS connection
>>>     // with a non-empty protocol name and TLSNextProto contains a
>>>     // map entry for that key (such as "h2"), then the func is
>>>     // called with the request's authority (such as "example.com"
>>>     // or "example.com:1234") and the TLS connection. The function
>>>     // must return a RoundTripper that then handles the request.
>>>     // If TLSNextProto is not nil, HTTP/2 support is not enabled
>>>     // automatically.
>>>     TLSNextProto map[string]func(authority string, c *tls.Conn) RoundTripper
>>>
>>>     // ProxyConnectHeader optionally specifies headers to send to
>>>     // proxies during CONNECT requests.
>>>     ProxyConnectHeader Header
>>>
>>>     // MaxResponseHeaderBytes specifies a limit on how many
>>>     // response bytes are allowed in the server's response
>>>     // header.
>>>     //
>>>     // Zero means to use a default limit.
>>>     MaxResponseHeaderBytes int64
>>>
>>>     // WriteBufferSize specifies the size of the write buffer used
>>>     // when writing to the transport.
>>>     // If zero, a default (currently 4KB) is used.
>>>     WriteBufferSize int
>>>
>>>     // ReadBufferSize specifies the size of the read buffer used
>>>     // when reading from the transport.
>>>     // If zero, a default (currently 4KB) is used.
>>>     ReadBufferSize int
>>>
>>>     // nextProtoOnce guards initialization of TLSNextProto and
>>>     // h2transport (via onceSetNextProtoDefaults)
>>>     nextProtoOnce      sync.Once
>>>     h2transport        h2Transport // non-nil if http2 wired up
>>>     tlsNextProtoWasNil bool        // whether TLSNextProto was nil when the Once fired
>>>
>>>     // ForceAttemptHTTP2 controls whether HTTP/2 is enabled when a non-zero
>>>     // Dial, DialTLS, or DialContext func or TLSClientConfig is provided.
>>>     // By default, use of any those fields conservatively disables HTTP/2.
>>>     // To use a custom dialer or TLS config and still attempt HTTP/2
>>>     // upgrades, set this to true.
>>>     ForceAttemptHTTP2 bool
>>> }```
---

> ## _TCP/IP에 대해_
>> TCP/IP란 인터넷을 구성하는 프로토콜 집합니다. 이 이름은 가장 널리 알려진 프로토콜인 TCP와 IP에서 따온 것이다.
>> TCP란 Transmission Control Protocol(전송 제어 프로토콜)의 줄임말이다. TCP 소프트웨어는 데이터를
>> 세그먼트(Segment)란 단위로 전송하는데 이를 TCP 패킷(Packet)이라 부른다. TCP의 두드러진 특성은 신뢰성을
>> 보장한다는 것이다. 다시 말해 프로그래머가 코드르 추가로 작성하지 않아도 패킷이 제대로 전달되도록 보장해준다.
>> 패킷이 전달됐다고 확인되지 않으면 그 패킷을 다시 보낸다. 특히 TCP 패킷은 연결을 생성하고 데이터를 전송하고 
>> 확인 메시지를 보내고 연결을 끊는데 사용된다.
>>
>> 두 머신 사이에 TCP 연결이 생성되면 전화가 연결될 때처럼 '풀 듀츨렉스(Full Duplex, 전이중)' 방식의
>> 가상 회로가 생성된다. 그런 다음 두 머신은 주고 받는 데이터가 정확하다는 것을 보장받는 상태로 지속적으로 
>> 통신할 수 있다. 어떤 이유로 연결이 끊기면 이러한 사실을 감지해서 애플리케이션에게 알려준다.
>>
>> IP는 Internet Protocol(인터넷 프로토콜)의 줄임말이다. IP의 가장 두드러진 특징은 자체적으로 신뢰성을
>> 보장하지 않는다는 것이다. IP는 TCP/IP 네트워크에 돌아다닐 데이터를 캡슐화한다. 출발지(Source)
>> 호스트에서 IP 주소로 지정한 목적지(Destination) 호스트로 패킷을 전달해야 하기 때문이다. IP 패킷의
>> 경로를 찾는 라우터(Router)라는 전용 전용 장비가 있지만, TCP/IP를 지원하는 장치라면 모두 기본적인 라우팅
>> 기능은 수행한다.
>>
>> UDP(User Datagram Protocol)는 IP를 기반으로 작동하며 신뢰성을 보장하지 않는다. 이처럼 신뢰성 보장
>> 기능이 빠져 있기 때문에 UDP는 TCP보다 간편하다. 따라서 UDP 메시지는 도중에 사라질 수도 있고 중복될 수도
>> 있고 보내는 순서와 달리 도착할 수 있다. 게다가 수신자가 처리하는 속도보다 빠르게 전달될 수도 있다. 그래서
>> UDP는 안정성보다 속도가 중요할 때 주로 사용한다.

---

> ## _IPv4와 IPv6에 대해_
>> IP의 첫 번쨰 버전을 IPv4라 부르는데 이는 그 뒤에 나온 버전인 IPv6와 구분하기 위해서다.
>>
>> IPv4가 가진 주된 문제점은 IP 주소가 거의 고갈될 수준에 이르렀다는 것이다. IPv6를 개발한 주된 이유가 바로
>> 여기에 있다. IPv4에서는 주소를 32비트로만 표현해서 총 4,294,967,296개의 IP 주소만 가질 수 있다. 반면
>> IPv6는 128비트로 주소를 표현한다.
>>
>> IPv4 주소는 10.20.32.245와 같은 형식으로 표현한다. 반면 IPv6에서는 3fce:1706:4523:3:150:f8ff:f221:56cf와
>> 같은 형식으로 표현한다.
 
---

> ## _nc(1) 커맨드라인 유틸리티_
>> nc(1) 유틸리티는 netcat(1)이라고도 부르며 TCP/IP 서버와 클라이언트를 테스트하는데 굉장히 유용하다.
>> nc(1)을 TCP 서비스의 클라이언트로 사용할 수 있다. 예를 들어, TCP 서버가 10.10.1.123이란 IP 주소를
>> 갖고 1234번 포트를 사용하는 머신에서 제공된다면 다음과 같이 실행할 수 있다.
>> ```
>> $ nc 10.10.1.123 1234
>> ```
>> nc(1)의 기본 설정에 따르면 TCP 프로토콜을 사용한다. 하지만 nc(1)에
>>
>> -u 플래그를 지정하면 UDP를 사용하게할 수 있다.
>>
>> -l 옵션을 지정하면 netcat(1)을 서버로 구동할 수 있다. 다시 말해 netcat(1)에 지정한 포트 번호로부터
>> 연결이 들어오길 기다린다.
>>
>> -v와 -vv 옵션을 지정하면 출력을 상세하게 표현한다. 네트워크 연결에 관련된 장애를 해결할 때 유용한 기능이다.
>>
>> netcat(1)를 HTTP 애플리케이션을 테스트하는 용도에 그치지 않고 TCP 및 UDP 클라이언트와 서버를 구현할 때도
>> 중요한 역할을 한다.