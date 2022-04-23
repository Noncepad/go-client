package basic

import (
	"errors"

	pbt "github.com/noncepad/client/proto/auth"
)

type RefreshArgs struct {
	ApiKey string
}
type RefreshResults struct {
	Token  string
	Expire int64
}

func (e1 Basic) Refresh(args RefreshArgs, results *RefreshResults) error {

	ans, err := e1.sessionClient.CreateByApiKey(e1.ctx, &pbt.ApiKeySessionRequest{Key: args.ApiKey})
	if err != nil {
		return err
	}

	if ans.Detail == nil {
		return errors.New("no detail")
	}

	*results = RefreshResults{Token: ans.Jwt, Expire: ans.Detail.Expire}

	return nil
}
