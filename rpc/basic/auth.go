package basic

import (
	"context"
	"errors"
	"log"
	"time"

	pbt "github.com/noncepad/client/proto/auth"
)

func (in *internal) look_up_by_api_key() *pbt.ApiKeySessionResponse {
	if in.session == nil {
		err := in.refresh()
		if err != nil {
			return nil
		}
	}

	return in.session
}

func (in *internal) refresh() error {

	ans, err := in.sessionClient.CreateByApiKey(in.ctx, &pbt.ApiKeySessionRequest{Key: in.key})
	if err != nil {
		log.Print(err)
		return err
	} else {
		if ans.Detail == nil {
			log.Print("serverside error")
			return errors.New("serverside error")
		}
		t := time.Unix(ans.Detail.Expire, 0)
		if t.Before(time.Now()) {
			log.Print("session expired")
			return errors.New("session expired from server")
		}
		in.session = ans
		go loopDelete(in.ctx, t, in.deleteC)
		return nil
	}
}

func loopDelete(ctx context.Context, expire time.Time, deleteC chan<- struct{}) {
	timeC := time.After(time.Until(expire))
	doneC := ctx.Done()
	select {
	case <-timeC:
		deleteC <- struct{}{}
	case <-doneC:
	}
}
