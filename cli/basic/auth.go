package basic

import (
	"errors"

	pbt "github.com/noncepad/client/proto/auth"
)

func (e1 *external) refresh(args []string) error {

	s := pbt.NewSessionClient(e1.conn)
	ans, err := s.CreateByApiKey(e1.ctx, &pbt.ApiKeySessionRequest{Key: e1.config.ApiKey})
	if err != nil {
		return err
	}

	if ans.Detail == nil {
		return errors.New("no detail")
	}

	e1.session = ans

	err = e1.SaveSession()
	if err != nil {
		return err
	}

	return nil
}
