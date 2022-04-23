package rpc

import (
	"context"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"time"

	basic "github.com/noncepad/client/rpc/basic"
	test "github.com/noncepad/client/rpc/test"
)

type external struct {
}

func Run(ctx context.Context, config *basic.Configuration, serverErrorC chan<- error) error {
	server := rpc.NewServer()

	err := server.Register(new(test.Arith))
	if err != nil {
		log.Print(err)
		return err
	}

	b, err := basic.Create(ctx, config)
	if err != nil {
		log.Print(err)
		return err
	}
	err = server.Register(b)
	if err != nil {
		log.Print(err)
		return err
	}

	if err != nil {
		log.Print(err)
		return err
	}

	s := new(Server)
	s.Rpc = server

	server.HandleHTTP("/jsonrpc", "/jsonrpc_debug")

	httpServer := &http.Server{
		Addr:        config.ListenUrl,
		Handler:     s,
		ReadTimeout: 5 * time.Second,
	}

	go loopServerListen(httpServer, serverErrorC)
	go loopClose(ctx, httpServer)
	log.Print("finished start up with no error")

	return nil
}

type Server struct {
	Rpc *rpc.Server
}

type HttpConn struct {
	in  io.Reader
	out io.Writer
}

func (c *HttpConn) Read(p []byte) (n int, err error)  { return c.in.Read(p) }
func (c *HttpConn) Write(d []byte) (n int, err error) { return c.out.Write(d) }
func (c *HttpConn) Close() error                      { return nil }

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/jsonrpc" {
		serverCodec := jsonrpc.NewServerCodec(&HttpConn{in: r.Body, out: w})
		w.Header().Set("Content-type", "application/json")
		w.WriteHeader(200)
		err := s.Rpc.ServeRequest(serverCodec)
		if err != nil {
			log.Printf("Error while serving JSON request: %v", err)
			http.Error(w, "Error while serving JSON request, details have been logged.", 500)
			return
		}
	}
}

func loopServerListen(server *http.Server, errorC chan<- error) {
	err := server.ListenAndServe()
	if err != nil {
		log.Print(err)
		errorC <- err
	}
}

func loopClose(ctx context.Context, l *http.Server) {
	<-ctx.Done()
	l.Close()
}
