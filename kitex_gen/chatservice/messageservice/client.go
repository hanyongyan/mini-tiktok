// Code generated by Kitex v0.4.4. DO NOT EDIT.

package messageservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	chatservice "mini_tiktok/kitex_gen/chatservice"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	MessageAction(ctx context.Context, Req *chatservice.MessageActionReq, callOptions ...callopt.Option) (r *chatservice.MessageActionResp, err error)
	MessageChat(ctx context.Context, Req *chatservice.MessageChatReq, callOptions ...callopt.Option) (r *chatservice.MessageChatResp, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kMessageServiceClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kMessageServiceClient struct {
	*kClient
}

func (p *kMessageServiceClient) MessageAction(ctx context.Context, Req *chatservice.MessageActionReq, callOptions ...callopt.Option) (r *chatservice.MessageActionResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageAction(ctx, Req)
}

func (p *kMessageServiceClient) MessageChat(ctx context.Context, Req *chatservice.MessageChatReq, callOptions ...callopt.Option) (r *chatservice.MessageChatResp, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.MessageChat(ctx, Req)
}
