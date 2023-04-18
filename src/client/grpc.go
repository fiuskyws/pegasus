package client

import (
	"context"

	"github.com/fiuskyws/pegasus/src/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type (
	gRPC struct {
		conn   *grpc.ClientConn
		client proto.PegasusClient
	}
)

func NewGRPC(target string) Client {
	conn, err := grpc.Dial(target)
	if err != nil {
		zap.L().Panic(err.Error())
	}
	client := proto.NewPegasusClient(conn)
	return &gRPC{
		conn:   conn,
		client: client,
	}
}

func (g *gRPC) Ping() error {
	_, err := g.client.GetTopics(context.Background(), &proto.GetTopicsRequest{})
	return err
}

func (g *gRPC) Connect() {}

func (g *gRPC) Close() error {
	return g.conn.Close()
}
