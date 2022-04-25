package basic

import (
	"context"
	"errors"
	"log"
	"time"

	pbt "github.com/noncepad/client/proto/auth"
)

type GetJwtArgs struct {
}

type GetJwtResponse struct {
	Jwt string `json:"token"`
}

func (e1 Basic) GetJwt(args GetJwtArgs, results *GetJwtResponse) error {
	ans := e1.look_up_by_api_key()
	if ans == nil {
		return errors.New("no session token")
	}
	*results = GetJwtResponse{Jwt: ans.GetJwt()}
	return nil
}

func (e1 Basic) look_up_by_api_key() *pbt.ApiKeySessionResponse {
	ansC := make(chan *pbt.ApiKeySessionResponse, 1)
	e1.internalC <- func(in *internal) {
		ansC <- in.look_up_by_api_key()
	}
	return <-ansC
}

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
