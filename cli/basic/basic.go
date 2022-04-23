package basic

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"

	"github.com/noncepad/client/cli"
	pbt "github.com/noncepad/client/proto/auth"
	pbsaas "github.com/noncepad/client/proto/saas"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type external struct {
	config          cli.Configuration
	ctx             context.Context
	conn            *grpc.ClientConn
	session         *pbt.ApiKeySessionResponse
	version         string
	serviceToString map[pbsaas.Service]string
	serviceToEnum   map[string]pbsaas.Service
}

func Create(ctx context.Context, config *cli.Configuration) (cli.Client, error) {

	// check for jwt
	e1 := &external{conn: nil, config: *config, ctx: ctx, session: nil}

	return e1, nil
}

func (e1 *external) init() error {
	if e1.conn != nil {
		return nil
	}

	e1.init_service()

	var err error
	e1.conn, err = grpc.DialContext(e1.ctx, fmt.Sprintf("%s:%d", e1.config.Host, e1.config.Port), grpc.WithInsecure())
	if err != nil {
		return err
	}
	// make sure we have a session token

	err = e1.LoadSession()
	if err != nil {
		err = e1.refresh([]string{})
		if err != nil {
			return err
		}
		err = e1.SaveSession()
		if err != nil {
			return err
		}
	}

	return nil
}

func (e1 *external) jwt_file_path() string {
	return fmt.Sprintf("%s/.solmate.json", os.Getenv("HOME"))
}

func (e1 *external) LoadSession() error {

	file, err := os.Open(e1.jwt_file_path())
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}
	e1.session = new(pbt.ApiKeySessionResponse)
	err = json.Unmarshal(data, e1.session)
	if err != nil {
		return err
	}

	if e1.session.Detail.Expire < time.Now().Unix() {
		e1.session = nil
		os.Remove(e1.jwt_file_path())
		return errors.New("expired token")
	}

	return nil
}

func (e1 *external) Ctx() context.Context {
	// e1.session must be filled in first
	return SetCtx(e1.ctx, e1.session.Jwt)
}

func SetCtx(ctx context.Context, token string) context.Context {
	md := metadata.New(map[string]string{"Authorization": token})
	return metadata.NewOutgoingContext(ctx, md)
}

func (e1 *external) SaveSession() error {
	if e1.session == nil {
		return errors.New("no session")
	}
	data, err := json.Marshal(e1.session)
	if err != nil {
		return err
	}
	path := e1.jwt_file_path()
	os.Remove(path)
	new, err := os.Create(path)
	if err != nil {
		return err
	}
	defer new.Close()
	err = os.Chmod(path, 0600)
	if err != nil {
		return err
	}

	t := len(data)
	i := 0
	for i < t {
		n, err := new.Write(data[i:])
		if err != nil {
			return err
		}
		i += n
	}
	return nil
}

type Command = string

const (
	CMD_VALIDATOR Command = "validator"
	CMD_SERVICE   Command = "service"
	CMD_DEPOSIT   Command = "deposit"
	CMD_REFRESH   Command = "refresh"
	CMD_VERSION   Command = "version"
	CMD_INFO      Command = "info"
)

func (e1 *external) Run(args []string) error {

	if len(args) < 1 {
		return errors.New("insufficient arguments")
	}

	cmd := args[0]
	var err error

	// let the user figure out his/her version for debugging purposes before attempting to connect to the server
	if cmd == CMD_VERSION {
		err = e1.Version()
		return err
	}

	err = e1.init()
	if err != nil {
		return err
	}

	switch cmd {
	case CMD_VALIDATOR:
		err = e1.validator(args[1:])
	case CMD_SERVICE:
		err = e1.service(args[1:])
	case CMD_DEPOSIT:
		err = e1.deposit(args[1:])
	case CMD_REFRESH:
		err = e1.refresh(args[1:])
	case CMD_INFO:
		err = e1.info()
	default:
		err = errors.New("no such command exists")
		e1.Help()
	}
	if err != nil {
		return err
	}
	return nil
}

func (e1 *external) Help() {
	fmt.Print(`
Commands
=============

validator - create and destory validator instances

service - subscribe and subscribe to services

deposit -  add funds (SOL or USDC) to your account

refresh - use the API key and retrieve a new session token
	`)
}

func (e1 *external) print_str(a string) {
	fmt.Printf("%s\n", a)
}

func (e1 *external) print_json(a interface{}) {
	data, err := json.Marshal(a)
	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Printf("%s\n", data)
	}
}
