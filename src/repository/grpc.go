package repository

import (
	"context"
	"time"

	"github.com/fiuskyws/pegasus/src/manager"
	"github.com/fiuskyws/pegasus/src/message"
	"github.com/fiuskyws/pegasus/src/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GRPCRepo struct {
	proto.UnimplementedPegasusServer
	mgr *manager.Manager
}

func NewGRPCRepo(mgr *manager.Manager) *GRPCRepo {
	return &GRPCRepo{
		mgr: mgr,
	}
}

// CreateTopic -
func (g *GRPCRepo) CreateTopic(ctx context.Context, req *proto.CreateTopicRequest) (*proto.CreateTopicResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	errChan := make(chan error, 1)

	go func() {
		errChan <- g.mgr.NewTopic(req.Name)
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		if err != nil {
			return nil, status.Error(codes.Canceled, err.Error())
		}
		return &proto.CreateTopicResponse{
			Error: "",
		}, nil
	}
}

// GetTopics returns a list of all topics.
func (g *GRPCRepo) GetTopics(ctx context.Context, _ *proto.GetTopicsRequest) (*proto.GetTopicsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	topics := make(chan []string, 1)

	go func() {
		topics <- g.mgr.GetTopicNames()
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case tpcs := <-topics:
		return &proto.GetTopicsResponse{
			Topics: tpcs,
		}, nil
	}
}

// Send - Endpoint for inserting messages into a Topic.
func (g *GRPCRepo) Send(ctx context.Context, req *proto.SendRequest) (*proto.SendResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	msgChan := make(chan *message.Message, 1)
	errChan := make(chan error, 1)
	go func() {
		msg, err := message.FromRequest(req)
		if err != nil {
			errChan <- err
			return
		}
		if err = g.mgr.Send(msg); err != nil {
			zap.L().Error(err.Error())
			errChan <- err
			return
		}
		msgChan <- msg
	}()
	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		zap.L().Error(err.Error())
		return nil, status.Error(codes.Canceled, err.Error())
	case <-msgChan:
		return &proto.SendResponse{
			Message: "message sent!",
		}, nil
	}
}

// Pop - Retrieves and delete the first item in the Topic's queue.
func (g *GRPCRepo) Pop(ctx context.Context, req *proto.PopRequest) (*proto.PopResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	msgChan := make(chan *message.Message, 1)
	errChan := make(chan error, 1)

	go func() {
		msg, err := g.mgr.Pop(req.TopicName)
		if err != nil {
			errChan <- err
			return
		}
		msgChan <- msg
	}()

	select {
	case <-ctx.Done():
		return nil, status.Error(codes.DeadlineExceeded, "timeout")
	case err := <-errChan:
		zap.L().Error(err.Error())
		return nil, status.Error(codes.Canceled, err.Error())
	case msg := <-msgChan:
		return &proto.PopResponse{
			TopicName: msg.TopicName,
			Body:      string(msg.Body),
		}, nil
	}
}

// Consumer -
func (g *GRPCRepo) Consumer(req *proto.ConsumerRequest, srv proto.Pegasus_ConsumerServer) error {
	msgChan := make(chan *message.Message, 1)
	errChan := make(chan error, 1)

	go func() {
		exit := false
		for !exit {
			msg, err := g.mgr.Pop(req.TopicName)
			if err != nil {
				errChan <- err
				exit = true
				return
			}
			msgChan <- msg
		}
	}()

	for {
		select {
		case <-srv.Context().Done():
			if err := srv.Context().Err(); err != nil {
				zap.L().Error(err.Error())
				return err
			}
			return nil
		case err := <-errChan:
			zap.L().Error(err.Error())
			return err
		case msg := <-msgChan:
			resp := proto.ConsumerResponse{
				TopicName: msg.TopicName,
				Body:      string(msg.Body),
			}
			if err := srv.Send(&resp); err != nil {
				zap.L().Error(err.Error())
				return err
			}
		}
	}
}

// Producer -
func (g *GRPCRepo) Producer(srv proto.Pegasus_ProducerServer) error {
	// TODO: Encapsulate this on it's own Job.
	exit := false
	done := srv.Context().Done()
	req, err := srv.Recv()
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	topic, err := g.mgr.GetTopic(req.TopicName)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	msg, err := message.FromRequest(req)
	if err != nil {
		zap.L().Error(err.Error())
		return err
	}

	if err := topic.Send(msg); err != nil {
		zap.L().Error(err.Error())
		return err
	}

	for !exit {
		select {
		case <-done:
			if err := srv.Context().Err(); err != nil {
				zap.L().Error(err.Error())
				return err
			}
			return nil
		default:
			req, err := srv.Recv()
			if err != nil {
				zap.L().Error(err.Error())
				return err
			}

			msg, err := message.FromRequest(req)
			if err != nil {
				zap.L().Error(err.Error())
				return err
			}

			if err := topic.Send(msg); err != nil {
				zap.L().Error(err.Error())
				return err
			}
		}
	}
	return nil
}
