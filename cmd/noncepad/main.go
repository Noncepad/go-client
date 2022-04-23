package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/noncepad/client/cli"
	"github.com/noncepad/client/cli/basic"
	log "github.com/sirupsen/logrus"
)

func main() {
	args := os.Args
	log.Infof("All arguments: %v\n", args)

	argsWithoutProgram := os.Args[1:]
	log.Errorf("Arguments without program name: %v\n", argsWithoutProgram)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sigChan := make(chan os.Signal, 1)
	go loopSignal(sigChan, cancel)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	config, err := cli.ConfigFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	client, err := basic.Create(ctx, config)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Run(args)
	if err != nil {
		fmt.Print(err)
		fmt.Print("\n")
		os.Exit(1)
	}

}

func loopSignal(sigChan <-chan os.Signal, cancel context.CancelFunc) {
	<-sigChan
	log.Print("received signal for the server to shutdown")
	cancel()
}
