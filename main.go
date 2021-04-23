package main

import (
	"fmt"
	"net/http"

	"github.com/EDDYCJY/go-gin-example/pkg/setting"
	"github.com/EDDYCJY/go-gin-example/routers"
)

func main() {
	// Default returns an Engine instance with the Logger and Recovery middleware already attached.
	// 得到一个默认的路由router的上下文，通过这个上下文可以做一些路由的设置
	// gin.Context主要能够处理url,处理json
	//

	// router逻辑写在main函数会显得逻辑混乱（是一块独立的逻辑），因此抽取出来放在一个单独的包下
	// router := gin.Default()
	// router.GET("/test", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "test",
	// 	})
	// })

	router := routers.InitRouter()

	//gin.H就是不用自己写的一个K,V
	// http包定义的一个结构体，定义server的一些信息
	// 这里用指针的原因查看ListenAndServe()源码可以发现定义的时候传的就是http包下的Server结构体指针
	//因此这里要用指针，而不是传递副本
	s := &http.Server{ // %d 十进制的整数，Sprintf 作用是按要求格式字符串
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router, //gin.engine 是作为一个handler传给net的http包的
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20, // 512k
	}
	//其中包含了在gin.engine里设置的router信息
	err := s.ListenAndServe() // 只是在server的8000端口建立起了一个监听
	if err != nil {
		fmt.Println(err)
	}
}

/*

func (srv *Server) ListenAndServe() error {
    addr := srv.Addr
    if addr == "" {
        addr = ":http"
    }
    ln, err := net.Listen("tcp", addr)
    if err != nil {
        return err
    }
    return srv.Serve(tcpKeepAliveListener{ln.(*net.TCPListener)})
}

*/

//  http包实现了server监听端口以及其他设置
// gin包实现了路由
/*
:A Server defines parameters for running an HTTP server. The zero value for Server is a valid configuration.

type Server struct {
    // Addr optionally specifies the TCP address for the server to listen on,
    // in the form "host:port". If empty, ":http" (port 80) is used.
    // The service names are defined in RFC 6335 and assigned by IANA.
    // See net.Dial for details of the address format.
    Addr string

    Handler Handler // handler to invoke, http.DefaultServeMux if nil

    // TLSConfig optionally provides a TLS configuration for use
    // by ServeTLS and ListenAndServeTLS. Note that this value is
    // cloned by ServeTLS and ListenAndServeTLS, so it's not
    // possible to modify the configuration with methods like
    // tls.Config.SetSessionTicketKeys. To use
    // SetSessionTicketKeys, use Server.Serve with a TLS Listener
    // instead.
    TLSConfig *tls.Config

    // ReadTimeout is the maximum duration for reading the entire
    // request, including the body.
    //
    // Because ReadTimeout does not let Handlers make per-request
    // decisions on each request body's acceptable deadline or
    // upload rate, most users will prefer to use
    // ReadHeaderTimeout. It is valid to use them both.
    ReadTimeout time.Duration

    // ReadHeaderTimeout is the amount of time allowed to read
    // request headers. The connection's read deadline is reset
    // after reading the headers and the Handler can decide what
    // is considered too slow for the body. If ReadHeaderTimeout
    // is zero, the value of ReadTimeout is used. If both are
    // zero, there is no timeout.
    ReadHeaderTimeout time.Duration

    // WriteTimeout is the maximum duration before timing out
    // writes of the response. It is reset whenever a new
    // request's header is read. Like ReadTimeout, it does not
    // let Handlers make decisions on a per-request basis.
    WriteTimeout time.Duration

    // IdleTimeout is the maximum amount of time to wait for the
    // next request when keep-alives are enabled. If IdleTimeout
    // is zero, the value of ReadTimeout is used. If both are
    // zero, there is no timeout.
    IdleTimeout time.Duration

    // MaxHeaderBytes controls the maximum number of bytes the
    // server will read parsing the request header's keys and
    // values, including the request line. It does not limit the
    // size of the request body.
    // If zero, DefaultMaxHeaderBytes is used.
    MaxHeaderBytes int

    // TLSNextProto optionally specifies a function to take over
    // ownership of the provided TLS connection when an ALPN
    // protocol upgrade has occurred. The map key is the protocol
    // name negotiated. The Handler argument should be used to
    // handle HTTP requests and will initialize the Request's TLS
    // and RemoteAddr if not already set. The connection is
    // automatically closed when the function returns.
    // If TLSNextProto is not nil, HTTP/2 support is not enabled
    // automatically.
    TLSNextProto map[string]func(*Server, *tls.Conn, Handler)

    // ConnState specifies an optional callback function that is
    // called when a client connection changes state. See the
    // ConnState type and associated constants for details.
    ConnState func(net.Conn, ConnState)

    // ErrorLog specifies an optional logger for errors accepting
    // connections, unexpected behavior from handlers, and
    // underlying FileSystem errors.
    // If nil, logging is done via the log package's standard logger.
    ErrorLog *log.Logger

    // BaseContext optionally specifies a function that returns
    // the base context for incoming requests on this server.
    // The provided Listener is the specific Listener that's
    // about to start accepting requests.
    // If BaseContext is nil, the default is context.Background().
    // If non-nil, it must return a non-nil context.
    BaseContext func(net.Listener) context.Context

    // ConnContext optionally specifies a function that modifies
    // the context used for a new connection c. The provided ctx
    // is derived from the base context and has a ServerContextKey
    // value.
    ConnContext func(ctx context.Context, c net.Conn) context.Context

    inShutdown atomicBool // true when when server is in shutdown

    disableKeepAlives int32     // accessed atomically.
    nextProtoOnce     sync.Once // guards setupHTTP2_* init
    nextProtoErr      error     // result of http2.ConfigureServer if used

    mu         sync.Mutex
    listeners  map[*net.Listener]struct{}
    activeConn map[*conn]struct{}
    doneChan   chan struct{}
    onShutdown []func()
}

*/
