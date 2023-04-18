package client

type (
	Client interface {
		Connect()
		Close() error
		Ping() error
	}
)
