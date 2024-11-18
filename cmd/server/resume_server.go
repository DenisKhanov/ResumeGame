package main

import (
	"context"
	"github.com/DenisKhanov/ResumeGame/internal/app/server"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	gameServer, err := server.NewApp(ctx)
	if err != nil {
		logrus.Fatalf("failed to init server app: %s", err.Error())
	}
	if gameServer != nil {
		gameServer.Run()
	} else {
		logrus.Fatal("failed to run server app")
	}

}
