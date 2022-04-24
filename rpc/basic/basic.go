package basic

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"os"
	"strconv"

	b1 "github.com/noncepad/client/cli/basic"
	pbt "github.com/noncepad/client/proto/auth"
	pbc "github.com/noncepad/client/proto/customer"
	pbjob "github.com/noncepad/client/proto/job"
	pbsaas "github.com/noncepad/client/proto/saas"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type Configuration struct {
	UseSSL    bool   `json:"ssl"`
	Host      string `json:"host"`
	Port      uint16 `json:"port"`
	Cacert    string `json:"cacert,omitempty"`
	HostName  string `json:"hostname,omitempty"`
	ListenUrl string `json:"listen"`
	ApiKey    string `json:"api_key"`
}

func ConfigDefaultHostPort(c *Configuration) {
	host := "grpc.test.noncepad.com"
	c.Host = host
	c.HostName = host
	c.Port = 443
	c.UseSSL = true
	c.ListenUrl = "0.0.0.0:8080"
}

func ConfigFromEnv() (*Configuration, error) {
	c := new(Configuration)
	var present bool

	c.ListenUrl, present = os.LookupEnv("LISTEN_URL")
	if !present {
		c.ListenUrl = "0.0.0.0:8080"
	}

	c.Host, present = os.LookupEnv("SOLMATE_HOST")
	if !present {
		ConfigDefaultHostPort(c)
	} else {
		portStr, present := os.LookupEnv("SOLMATE_PORT")
		if !present {
			return nil, errors.New("no port")
		}
		portNum, err := strconv.ParseUint(portStr, 10, 16)
		if err != nil {
			return nil, err
		}
		c.Port = uint16(portNum)

		if os.Getenv("USE_SSL") == "1" {
			cacert, present := os.LookupEnv("CACERT")
			if present {
				c.Cacert = cacert
				c.UseSSL = true
			} else {
				c.Cacert = ""
				c.UseSSL = false
			}

			c.HostName, present = os.LookupEnv("SOLMATE_HOST_NAME")
			if !present {
				c.HostName = c.Host
			}
		}
	}

	c.ApiKey, present = os.LookupEnv("API_KEY")
	if !present {
		return nil, errors.New("no API_KEY")
	}

	return c, nil
}

type Basic struct {
	config          Configuration
	ctx             context.Context
	conn            *grpc.ClientConn
	version         string
	serviceToString map[pbsaas.Service]string
	serviceToEnum   map[string]pbsaas.Service
	internalC       chan<- func(*internal)
	sessionClient   pbt.SessionClient
	saasClient      pbc.SaasClient
	solanaClient    pbc.SolanaClient
	jobClient       pbjob.OwnerClient
}

func Dial(ctx context.Context, config *Configuration) (*grpc.ClientConn, error) {
	if config.UseSSL {
		return grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithTransportCredentials(credentials.NewTLS(&tls.Config{
			ServerName: config.HostName,
		})))
	} else {
		return grpc.DialContext(ctx, fmt.Sprintf("%s:%d", config.Host, config.Port), grpc.WithInsecure())
	}

}

func Create(ctx context.Context, config *Configuration) (Basic, error) {
	conn, err := Dial(ctx, config)
	if err != nil {
		return Basic{}, err
	}
	internalC := make(chan func(*internal), 10)
	a, b := init_service()
	// check for jwt
	e1 := Basic{
		conn:            conn,
		config:          *config,
		ctx:             ctx,
		serviceToString: a,
		serviceToEnum:   b,
		internalC:       internalC,
		sessionClient:   pbt.NewSessionClient(conn),
		saasClient:      pbc.NewSaasClient(conn),
		solanaClient:    pbc.NewSolanaClient(conn),
		jobClient:       pbjob.NewOwnerClient(conn),
	}

	go loopInternal(ctx, internalC, config.ApiKey, e1.sessionClient)

	return e1, nil
}

func (e1 Basic) Ctx() (context.Context, error) {
	sessionC := make(chan *pbt.ApiKeySessionResponse, 1)
	e1.internalC <- func(in *internal) {
		sessionC <- in.look_up_by_api_key()
	}
	session := <-sessionC
	if session == nil {
		return nil, errors.New("api key is not valid")
	}

	// e1.session must be filled in first
	return b1.SetCtx(e1.ctx, session.Jwt), nil
}
