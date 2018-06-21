package web

import (
	"bytes"
	"fmt"
	"game/common/server"
	"game/common/utils"
	"net"
	"net/http"
	"path"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
	"time"

	"golang.org/x/net/websocket"

	log "github.com/cihub/seelog"
)

type WebService struct {
	server.BaseService

	port          int
	acceptTimeout time.Duration
	listener      *net.TCPListener
	StaticDir     string
	routes        []route
}
type route struct {
	r           string
	cr          *regexp.Regexp
	method      string
	handler     reflect.Value
	httpHandler http.Handler
}

func NewWebService(port int, acceptTimeout time.Duration, staticDir string) *WebService {
	s := new(WebService)
	s.port = port
	s.acceptTimeout = acceptTimeout
	s.StaticDir = staticDir
	return s
}
func (s *WebService) Start() error {
	serverAddr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", s.port))
	if err != nil {
		return err
	}
	s.listener, err = net.ListenTCP("tcp", serverAddr)
	if err != nil {
		return err
	}
	log.Infof("listen tcp, port=%v", serverAddr)
	s.BaseService.Start()
	s.AsyncDo(func() {
		defer func() {
			recover()
			if s.listener != nil {
				log.Infof("defer web server close!")
				s.listener.Close()
				s.listener = nil
			}
		}()
		mux := http.NewServeMux()
		mux.Handle("/", s)
		err = http.Serve(s.listener, mux)
		if err != nil {
			return
		}
	})
	return nil

}
func (s *WebService) Close() {
	log.Infof("web server close!")
	s.BaseService.Close()
}
func (s *WebService) ServeHTTP(c http.ResponseWriter, req *http.Request) {
	s.Process(c, req)
}
func (s *WebService) Process(c http.ResponseWriter, req *http.Request) {
	route := s.routeHandler(req, c)
	if route != nil {
		route.httpHandler.ServeHTTP(c, req)
	}
}
func (s *WebService) routeHandler(req *http.Request, w http.ResponseWriter) (unused *route) {
	requestPath := req.URL.Path
	ctx := Context{w, req, map[string]string{}, s}
	ctx.SetHeader("webserver", "go", true)
	tm := time.Now().UTC()
	req.ParseForm()
	if len(req.Form) > 0 {
		for k, v := range req.Form {
			ctx.Params[k] = v[0]
		}
	}
	defer s.logRequest(ctx, tm)
	ctx.SetHeader("Date", utils.WebTime(tm), true)
	if req.Method == "GET" || req.Method == "HEAD" {
		if s.tryServingFile(requestPath, req, w) {
			return
		}
	}
	for i := 0; i < len(s.routes); i++ {
		route := s.routes[i]
		cr := route.cr
		if req.Method != route.method && !(req.Method == "HEAD" && route.method == "GET") {
			continue
		}
		if !cr.MatchString(requestPath) {
			continue
		}
		match := cr.FindStringSubmatch(requestPath)
		if len(match[0]) != len(requestPath) {
			continue
		}
		if route.httpHandler != nil {
			unused = &route
			return
		}
		ctx.SetHeader("Content-Type", "text/html; charset=utf-8", true)
		var args []reflect.Value
		handlerType := route.handler.Type()
		if requiresContext(handlerType) {
			args = append(args, reflect.ValueOf(&ctx))
		}
		for _, arg := range match[1:] {
			args = append(args, reflect.ValueOf(arg))
		}
		ret, err := s.safelyCall(route.handler, args)
		if err != nil {
			ctx.Abort(500, "Server Error")
		}
		if len(ret) == 0 {
			return
		}
		sval := ret[0]
		var content []byte
		if sval.Kind() == reflect.String {
			content = []byte(sval.String())
		} else if sval.Kind() == reflect.Slice && sval.Type().Elem().Kind() == reflect.Uint8 {
			content = sval.Interface().([]byte)
		}
		ctx.SetHeader("Content-Length", strconv.Itoa(len(content)), true)
		_, err = ctx.ResponseWriter.Write(content)
		if err != nil {
			log.Errorf("Error during write %v", err)
		}
		return
	}
	if req.Method == "GET" || req.Method == "HEAD" {
		if s.tryServingFile(path.Join(requestPath, "index.html"), req, w) {
			return
		} else if s.tryServingFile(path.Join(requestPath, "index.htm"), req, w) {
			return
		}
	}
	ctx.Abort(404, "Page not found")
	return

}
func (s *WebService) safelyCall(function reflect.Value, args []reflect.Value) (resp []reflect.Value, e interface{}) {
	defer func() {
		if err := recover(); err != nil {
			e = err
			resp = nil
			log.Infof("Handler crashed with error :%v", err)
			for i := 1; ; i++ {
				_, file, line, ok := runtime.Caller(i)
				if !ok {
					break
				}
				log.Infof(file, line)
			}
		}
	}()
	return function.Call(args), nil
}
func (s *WebService) tryServingFile(name string, req *http.Request, w http.ResponseWriter) bool {
	if s.StaticDir != "" {
		staticFile := path.Join(s.StaticDir, name)
		if utils.FileExists(staticFile) {
			http.ServeFile(w, req, staticFile)
			return true
		}
	}
	return false
}
func (s *WebService) logRequest(ctx Context, sTime time.Time) {
	req := ctx.Request
	requestPath := req.URL.Path
	duration := time.Now().Sub(sTime)
	var client string
	pos := strings.LastIndex(req.RemoteAddr, ":")
	if pos > 0 {
		client = req.RemoteAddr[0:pos]
	} else {
		client = req.RemoteAddr
	}
	var logEntry bytes.Buffer
	logEntry.WriteString(client)
	logEntry.WriteString("-" + req.Method + " " + requestPath)
	logEntry.WriteString("-" + duration.String())
	if len(ctx.Params) > 0 {
		logEntry.WriteString("-" + fmt.Sprintf("Params: %v\n", ctx.Params))
	}
	log.Info(logEntry.String())
}
func (s *WebService) addRoute(r string, method string, handler interface{}) {
	cr, err := regexp.Compile(r)
	if err != nil {
		log.Errorf("Error in route regex %q\n", r)
		return
	}
	switch handler.(type) {
	case http.Handler:
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, httpHandler: handler.(http.Handler)})
	case reflect.Value:
		fv := handler.(reflect.Value)
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, handler: fv})
	default:
		fv := reflect.ValueOf(handler)
		s.routes = append(s.routes, route{r: r, cr: cr, method: method, handler: fv})
	}
}

// Get adds a handler for the 'GET' http method for server s.
func (s *WebService) Get(route string, handler interface{}) {
	s.addRoute(route, "GET", handler)
}

// Post adds a handler for the 'POST' http method for server s.
func (s *WebService) Post(route string, handler interface{}) {
	s.addRoute(route, "POST", handler)
}

// Put adds a handler for the 'PUT' http method for server s.
func (s *WebService) Put(route string, handler interface{}) {
	s.addRoute(route, "PUT", handler)
}

// Delete adds a handler for the 'DELETE' http method for server s.
func (s *WebService) Delete(route string, handler interface{}) {
	s.addRoute(route, "DELETE", handler)
}

// Match adds a handler for an arbitrary http method for server s.
func (s *WebService) Match(method string, route string, handler interface{}) {
	s.addRoute(route, method, handler)
}

// Add a custom http.Handler. Will have no effect when running as FCGI or SCGI.
func (s *WebService) Handle(route string, method string, httpHandler http.Handler) {
	s.addRoute(route, method, httpHandler)
}

//Adds a handler for websockets. Only for webserver mode. Will have no effect when running as FCGI or SCGI.
func (s *WebService) Websocket(route string, httpHandler websocket.Handler) {
	s.addRoute(route, "GET", httpHandler)
}

func requiresContext(handlerType reflect.Type) bool {
	//if the method doesn't take arguments, no
	if handlerType.NumIn() == 0 {
		return false
	}

	//if the first argument is not a pointer, no
	a0 := handlerType.In(0)
	if a0.Kind() != reflect.Ptr {
		return false
	}
	//if the first argument is a context, yes
	if a0.Elem() == reflect.TypeOf(Context{}) {
		return true
	}

	return false
}
