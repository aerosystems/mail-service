package main

import (
	"context"
	"golang.org/x/sync/errgroup"
)

// @title Mail Service
// @version 1.0.1
// @description A part of microservice infrastructure, who responsible for sending emails

// @contact.name Artem Kostenko
// @contact.url https://github.com/aerosystems

// @license.name Apache 2.0
// @license.url https://www.apache.org/licenses/LICENSE-2.0.html

// @host gw.verifire.dev/mail
// @schemes https
// @BasePath /
func main() {
	app := InitApp()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	group, ctx := errgroup.WithContext(ctx)

	group.Go(func() error {
		return app.httpServer.Run()
	})

	group.Go(func() error {
		return app.rpcServer.Run()
	})

	group.Go(func() error {
		return app.handleSignals(ctx, cancel)
	})

	if err := group.Wait(); err != nil {
		app.log.Errorf("error occurred: %v", err)
	}
}
