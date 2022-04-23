package rpc

import (
	"context"
	"log"
	"net"
	"net/http"
	"net/rpc"

	basic "github.com/noncepad/client/rpc/basic"
	test "github.com/noncepad/client/rpc/test"
)

type external struct {
}

func Run(ctx context.Context, config *basic.Configuration, listenUrl string, serverErrorC chan<- error) error {

	err := rpc.Register(new(test.Arith))
	if err != nil {
		log.Print(err)
		return err
	}

	b, err := basic.Create(ctx, config)
	if err != nil {
		log.Print(err)
		return err
	}
	err = rpc.Register(b)
	if err != nil {
		log.Print(err)
		return err
	}

	if err != nil {
		log.Print(err)
		return err
	}
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", listenUrl)
	if err != nil {
		log.Print(err)
		return err
	}
	go loopServe(l, serverErrorC)
	go loopClose(ctx, l)
	log.Print("finished start up with no error")
	return nil
}

func loopClose(ctx context.Context, l net.Listener) {
	<-ctx.Done()
	l.Close()
}

func loopServe(l net.Listener, errorC chan<- error) {
	err := http.Serve(l, nil)
	errorC <- err
}
