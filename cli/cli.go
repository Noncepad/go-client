package cli

import (
	"encoding/json"
	"errors"
	"os"
	"strconv"
)

type Configuration struct {
	Host   string `json:"host"`
	Port   uint16 `json:"port"`
	ApiKey string `json:"key"`
}

func ConfigFromEnv() (*Configuration, error) {
	c := new(Configuration)
	var present bool
	var err error

	c.Host, present = os.LookupEnv("SOLMATE_HOST")
	if !present {
		return nil, errors.New("no host")
	}
	portStr, present := os.LookupEnv("SOLMATE_PORT")
	if !present {
		return nil, errors.New("no port")
	}
	portNum, err := strconv.ParseUint(portStr, 10, 16)
	if err != nil {
		return nil, err
	}
	c.Port = uint16(portNum)

	c.ApiKey, present = os.LookupEnv("API_KEY")
	if !present {
		return nil, errors.New("no api key")
	}

	return c, nil
}

func ConfigFromFile(filePath string) (*Configuration, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	c := new(Configuration)
	err = json.NewDecoder(file).Decode(c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

type Client interface {
	Run(args []string) error
}
