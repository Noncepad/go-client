package main_test

import (
	"context"
	"testing"

	"github.com/joho/godotenv"
	"github.com/noncepad/client/cli"
	"github.com/noncepad/client/cli/basic"
	log "github.com/sirupsen/logrus"
)

func TestRefresh(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"service", "menu", "--json"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}

	args = []string{"service", "portal"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}
}

func TestSubscribe(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"service", "sub", "validator_instance"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}

}

func TestCreateValidator(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"validator", "create"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}

}

func TestListValidator(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"validator", "ls"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}

}

func TestDestroyValidator(t *testing.T) {
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load("../../.env")
	if err != nil {
		t.Fatal(err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(func() {
		cancel()
	})
	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{"validator", "destroy", "fe1966fd-f38e-4bef-bcad-d768ffb174c8"}

	err = client.Run(args)
	if err != nil {
		log.Fatal(err)
	}

}
