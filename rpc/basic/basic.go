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
	UseSSL   bool   `json:"ssl"`
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	Cacert   string `json:"cacert,omitempty"`
	HostName string `json:"hostname,omitempty"`
}

func ConfigDefaultHostPort(c *Configuration) {
	host := "grpc.test.noncepad.com"
	c.Host = host
	c.HostName = host
	c.Port = 443
	c.UseSSL = true
}

func ConfigFromEnv() (*Configuration, error) {
	c := new(Configuration)
	var present bool

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

	return c, nil
}

type Basic struct {
	config          Configuration
	ctx             context.Context
	conn            *grpc.ClientConn
	version         string
	serviceToString map[pbsaas.Service]string
	serviceToEnum   map[string]pbsaas.Service

	sessionClient pbt.SessionClient
	saasClient    pbc.SaasClient
	solanaClient  pbc.SolanaClient
	jobClient     pbjob.OwnerClient
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

	a, b := init_service()
	// check for jwt
	e1 := Basic{
		conn:            conn,
		config:          *config,
		ctx:             ctx,
		serviceToString: a,
		serviceToEnum:   b,
		sessionClient:   pbt.NewSessionClient(conn),
		saasClient:      pbc.NewSaasClient(conn),
		solanaClient:    pbc.NewSolanaClient(conn),
		jobClient:       pbjob.NewOwnerClient(conn),
	}

	return e1, nil
}

type RequestWithJwt struct {
	Token string
}

func (e1 Basic) Ctx(token string) context.Context {
	// e1.session must be filled in first
	return b1.SetCtx(e1.ctx, token)
}
