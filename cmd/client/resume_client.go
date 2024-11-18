package main

import (
	"context"
	"github.com/DenisKhanov/ResumeGame/internal/app/client"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	gameClient, err := client.NewApp(ctx)
	if err != nil {
		logrus.Fatalf("failed to init client app: %s", err.Error())
	}
	if gameClient != nil {
		gameClient.Run()
	} else {
		logrus.Fatal("failed to run client app")
	}

}
