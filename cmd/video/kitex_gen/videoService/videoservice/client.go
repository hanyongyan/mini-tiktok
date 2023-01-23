// Code generated by Kitex v0.4.4. DO NOT EDIT.

package videoservice

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	videoService "mini_tiktok/cmd/video/kitex_gen/videoService"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	PublishAction(ctx context.Context, Req *videoService.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinPublishActionResponse, err error)
	Feed(ctx context.Context, Req *videoService.DouyinFeedRequest, callOptions ...callopt.Option) (r *videoService.DouyinFeedResponse, err error)
	PublishList(ctx context.Context, Req *videoService.DouyinPublishListRequest, callOptions ...callopt.Option) (r *videoService.DouyinPublishListResponse, err error)
	FavoriteAction(ctx context.Context, Req *videoService.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinFavoriteActionResponse, err error)
	FavoriteList(ctx context.Context, Req *videoService.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *videoService.DouyinFavoriteListResponse, err error)
	CommentAction(ctx context.Context, Req *videoService.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinCommentActionResponse, err error)
	CommentList(ctx context.Context, Req *videoService.DouyinCommentListRequest, callOptions ...callopt.Option) (r *videoService.DouyinCommentListResponse, err error)
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
	return &kVideoServiceClient{
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

type kVideoServiceClient struct {
	*kClient
}

func (p *kVideoServiceClient) PublishAction(ctx context.Context, Req *videoService.DouyinPublishActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinPublishActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishAction(ctx, Req)
}

func (p *kVideoServiceClient) Feed(ctx context.Context, Req *videoService.DouyinFeedRequest, callOptions ...callopt.Option) (r *videoService.DouyinFeedResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Feed(ctx, Req)
}

func (p *kVideoServiceClient) PublishList(ctx context.Context, Req *videoService.DouyinPublishListRequest, callOptions ...callopt.Option) (r *videoService.DouyinPublishListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PublishList(ctx, Req)
}

func (p *kVideoServiceClient) FavoriteAction(ctx context.Context, Req *videoService.DouyinFavoriteActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinFavoriteActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteAction(ctx, Req)
}

func (p *kVideoServiceClient) FavoriteList(ctx context.Context, Req *videoService.DouyinFavoriteListRequest, callOptions ...callopt.Option) (r *videoService.DouyinFavoriteListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.FavoriteList(ctx, Req)
}

func (p *kVideoServiceClient) CommentAction(ctx context.Context, Req *videoService.DouyinCommentActionRequest, callOptions ...callopt.Option) (r *videoService.DouyinCommentActionResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentAction(ctx, Req)
}

func (p *kVideoServiceClient) CommentList(ctx context.Context, Req *videoService.DouyinCommentListRequest, callOptions ...callopt.Option) (r *videoService.DouyinCommentListResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.CommentList(ctx, Req)
}