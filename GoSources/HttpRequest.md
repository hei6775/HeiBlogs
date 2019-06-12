# Golang 的 request 源码阅读

## Request 结构体

```golang
type Request struct {
	Method string
	URL *url.URL
	Proto      string // "HTTP/1.0"
	ProtoMajor int    // 1
	ProtoMinor int    // 0
	// Header contains the request header fields either received
	// by the server or to be sent by the client.
	//
	// If a server received a request with header lines,
	//
	//	Host: example.com
	//	accept-encoding: gzip, deflate
	//	Accept-Language: en-us
	//	fOO: Bar
	//	foo: two
	//
	// then
	//
	//	Header = map[string][]string{
	//		"Accept-Encoding": {"gzip, deflate"},
	//		"Accept-Language": {"en-us"},
	//		"Foo": {"Bar", "two"},
	//	}
	//
	// For incoming requests, the Host header is promoted to the
	// Request.Host field and removed from the Header map.
	//
	// HTTP defines that header names are case-insensitive. The
	// request parser implements this by using CanonicalHeaderKey,
	// making the first character and any characters following a
	// hyphen uppercase and the rest lowercase.
	//
	// For client requests, certain headers such as Content-Length
	// and Connection are automatically written when needed and
	// values in Header may be ignored. See the documentation
	// for the Request.Write method.
	Header Header

	// Body is the request's body.
	//
	// For client requests a nil body means the request has no
	// body, such as a GET request. The HTTP Client's Transport
	// is responsible for calling the Close method.
	//
	// For server requests the Request Body is always non-nil
	// but will return EOF immediately when no body is present.
	// The Server will close the request body. The ServeHTTP
	// Handler does not need to.
	Body io.ReadCloser
	GetBody func() (io.ReadCloser, error)
	ContentLength int64
	TransferEncoding []string
	Close bool
	Host string
	Form url.Values
	PostForm url.Values
	MultipartForm *multipart.Form
	Trailer Header
	RemoteAddr string
	RequestURI string
	// TLS allows HTTP servers and other software to record
	// information about the TLS connection on which the request
	// was received. This field is not filled in by ReadRequest.
	// The HTTP server in this package sets the field for
	// TLS-enabled connections before invoking a handler;
	// otherwise it leaves the field nil.
	// This field is ignored by the HTTP client.
	TLS *tls.ConnectionState
	Cancel <-chan struct{}
	Response *Response
	ctx context.Context
}
```

`request`的字段意思

- `Method`,`string`类型，显示 http 请求的方法（如 GET、POST 等），如果客户端传上来的是空值那么默认就是 GET 方法。
- `URL`,`*url.URL`类型，请求的`URL`
- `Proto`,`string`类型，`HTTP/1.0`
- `ProtoMajor`, `string`类型，`1`
- `ProtoMinor`, `string`类型，`0`

- `Hearder`, `Hearder`类型（`map[string][]string`），接受到的请求头信息解析到这个`map`中，`Host`字段会被删除，`key`不区分大小写（实际上会让所有的`key`变成第一个为大写字母，后面的所有都是小写字母）。`Content-Length`和`Connection`会被自动写入`Hearder`中。

- `Body`, `io.ReadCloser`类型（接口类型），客户端的`body`可能是空的（比如说`GET`方法），客户端传输的话就是调用`body`的`close`方法。而对于服务端而言，`body`永远都是非空的（`non-nil`），当当前没有`body`的时候会返回`EOF`。服务端会关闭这个`request.body`。

- `GetBody`, `func()(io.ReadCloser,error)`，GetBody 定义了一个可选的 `func`来返回 `Body` 的新副本。 当重定向需要多次读取主体时，它用于客户端请求。 使用 `GetBody` 仍然需要设置 `Body`。对于服务器请求，它未被使用。

- `ContentLength`,`int64`类型，记录关联内容的长度。 值`-1`表示长度未知。 值`> = 0`表示可以从 Body 读取给定的字节数。 对于客户端请求，具有非零主体的值为 0 也被视为未知。

- `TransferEncoding`,`[]string`类型。

- `CLose`,`bool`类型。`Close`表示在回复此请求后（对于服务器）或发送此请求并读取其响应（对于客户端）之后是否关闭连接。
  对于服务器请求，HTTP 服务器自动处理此请求，处理程序不需要此字段。
  对于客户端请求，设置此字段可防止在对相同主机的请求之间重复使用 TCP 连接，就像设置了`Transport.DisableKeepAlives`一样。

- `Host`, `string`类型。

- `PostForm`, `url.Values`类型，需要在主动调用`ParseForm`后才可以使用，包含 POST、PATCH、PUT 的请求 body 中的参数。

- `MultipartForm`, `*multipart.Form`类型，需要主动调用`ParseMultipartForm`后才可以读取，包含文件上传等表单信息。

- `Trailer` ,`Header`类型。
- `RemoteAddr`, `string`类型。
- `RequestURI`, `string`类型。
- `TLS`, `*tls.ConnectionState`类型。
- `Cancel`, `<-chan struct{}`类型。
- `Response`,`*Response`类型。
- `ctx` ,`context.Context`类型。

## Request 的方法

```golang
func (r *Request) Context() context.Context {
	if r.ctx != nil {
		return r.ctx
	}
	return context.Background()
}
```
