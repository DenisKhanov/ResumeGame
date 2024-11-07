package main

import (
	"context"
	"github.com/DenisKhanov/ResumeGame/internal/app/server"

	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()

	resumeGame, err := server.NewApp(ctx)
	if err != nil {
		logrus.Fatalf("failed to init app: %s", err.Error())
	}

	resumeGame.Run()
}
