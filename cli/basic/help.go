package basic

import (
	"errors"
	"fmt"
	"time"
)

func (e1 *external) Version() error {
	fmt.Printf("version=%s\n", e1.version)
	return nil
}

func (e1 *external) info() error {
	if e1.session == nil {
		return errors.New("no session")
	}
	fmt.Printf(`
Client Info

========

Version:  %s

User Id: %s

Session Expiry: %s
	`, e1.version, e1.session.Detail.UserId, time.Unix(e1.session.Detail.Expire, 0))

	return nil

}
