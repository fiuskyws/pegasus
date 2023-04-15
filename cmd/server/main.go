package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/fiuskyws/pegasus/src/manager"
	"github.com/fiuskyws/pegasus/src/server"
)

var (
	port   = uint(8090)
	topics = []string{
		"topic-1", "topic-2",
	}
)

func main() {

	exit := make(chan os.Signal, 1)

	signal.Notify(exit, syscall.SIGTERM, syscall.SIGINT)

	m := manager.NewManager()

	for _, topic := range topics {
		if err := m.NewTopic(topic); err != nil {
			panic(err)
		}
	}

	g := server.NewGRPC(m)

	if err := g.Start(port); err != nil {
		panic(err)
	}

	<-exit

	if err := g.Close(); err != nil {
		panic(err)
	}
}
