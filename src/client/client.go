package client

import "github.com/fiuskyws/pegasus/src/proto"

type (
	Client interface {
		Connect()
		Close() error
		Ping() error
		GetClient() proto.PegasusClient
	}
)
