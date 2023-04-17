package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fiuskyws/pegasus/src/manager"
	"github.com/fiuskyws/pegasus/src/server"
	"go.uber.org/zap"
)

var (
	port   = uint(8090)
	topics = []string{
		"topic-1", "topic-2",
	}
)

func main() {
	l, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	// TODO:
	// 	- Remove this implementation, MUST NOT use `ReplaceGlobals`
	undo := zap.ReplaceGlobals(l)
	defer undo()

	exit := make(chan os.Signal, 1)

	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	m := manager.NewManager()

	for _, topic := range topics {
		if err := m.NewTopic(topic); err != nil {
			zap.L().Panic(err.Error())
		}
	}

	g := server.NewGRPC(m)

	go func() {
		if err := g.Start(port); err != nil {
			zap.L().Panic(err.Error())
		}
	}()

	<-exit

	if err := g.Close(); err != nil {
		zap.L().Panic(err.Error())
	}
}
