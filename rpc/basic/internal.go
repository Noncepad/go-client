package basic

import (
	"context"
	"log"

	pbt "github.com/noncepad/client/proto/auth"
)

type ApiKey = string

type internal struct {
	ctx              context.Context
	sessionClient    pbt.SessionClient
	session          *pbt.ApiKeySessionResponse
	closeSignalCList []chan<- error
	key              ApiKey
	deleteC          chan<- struct{}
	jobId            int
	jobMap           map[int]*TxJob
}

func loopInternal(ctx context.Context, internalC <-chan func(*internal), apiKey ApiKey, sessionClient pbt.SessionClient) {
	var err error
	deleteC := make(chan struct{}, 1)

	in := new(internal)
	in.session = nil
	in.ctx = ctx
	in.closeSignalCList = make([]chan<- error, 0)
	in.sessionClient = sessionClient
	in.key = apiKey
	in.deleteC = deleteC
	in.jobMap = make(map[int]*TxJob)

	doneC := ctx.Done()
out:
	for {
		select {
		case <-doneC:
			break out
		case req := <-internalC:
			req(in)
		case <-deleteC:
			in.session = nil
		}
	}

	in.finish(err)
}

func (e1 Basic) CloseSignal() <-chan error {
	signalC := make(chan error, 1)
	e1.internalC <- func(in *internal) {
		in.closeSignalCList = append(in.closeSignalCList, signalC)
	}
	return signalC
}

func (in *internal) finish(err error) {
	log.Print(err)
	for i := 0; i < len(in.closeSignalCList); i++ {
		in.closeSignalCList[i] <- err
	}
}
