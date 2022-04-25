package rpc

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"net/url"
	"time"

	basic "github.com/noncepad/client/rpc/basic"
	test "github.com/noncepad/client/rpc/test"
)

type external struct {
}

func Run(ctx context.Context, config *basic.Configuration, serverErrorC chan<- error) error {
	server := rpc.NewServer()

	validatorUrl, err := url.Parse(config.ValidatorRpcUrl)
	if err != nil {
		return err
	}

	err = server.Register(new(test.Arith))
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
	s.validator = validatorUrl
	s.Rpc = server
	s.b = b

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
	validator *url.URL
	Rpc       *rpc.Server
	b         basic.Basic
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
	} else if r.Method == "POST" {
		jwtResults := new(basic.GetJwtResponse)

		err := s.b.GetJwt(basic.GetJwtArgs{}, jwtResults)
		if err != nil {
			//log.Println(err)
			http.Error(w, "Unauthorized API Key", 403)
			return
		}

		// proxy
		//adding the proxy settings to the Transport object
		//log.Printf("host=%s", s.validator.Host)
		client := &http.Client{}
		request, err := http.NewRequest("POST", fmt.Sprintf("https://%s/solana/validator", s.validator.Host), r.Body)
		if err != nil {
			log.Println(err)
			http.Error(w, "Error while proxying", 500)
			return
		}

		copyHeader(request.Header, r.Header)
		request.Header["JWT"] = []string{jwtResults.Jwt}
		request.Header["Cluster"] = []string{"mainnet-beta"}

		response, err := client.Do(request)
		if err != nil {
			//log.Println(err)
			http.Error(w, "Error while proxying", 500)
			return
		}
		defer response.Body.Close()
		copyHeader(w.Header(), response.Header)

		_, err = io.Copy(w, response.Body)
		if err != nil {
			//log.Println(err)
			http.Error(w, "Error while proxying", 500)
			return
		}
		//w.WriteHeader(response.StatusCode)

	}
}

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
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
