// Package hailuo provides the Hailuo video generation API client.
package hailuo

import (
	"context"

	"github.com/runapi-ai/core-sdk/go/base"
	"github.com/runapi-ai/core-sdk/go/core"
	"github.com/runapi-ai/core-sdk/go/option"
)

const (
	textToVideoPath  = "/api/v1/hailuo/text_to_video"
	imageToVideoPath = "/api/v1/hailuo/image_to_video"
)

// Client provides Hailuo text-to-video and image-to-video generation.
type Client struct {
	base.Base
	TextToVideo  *TextToVideo
	ImageToVideo *ImageToVideo
}

// NewClient creates a Hailuo client with the given options.
func NewClient(opts ...option.ClientOption) (*Client, error) {
	resolved, err := option.ResolveClientOptions(opts...)
	if err != nil {
		return nil, err
	}
	httpClient, err := core.NewHTTPClient(resolved)
	if err != nil {
		return nil, err
	}
	return NewClientWithHTTP(httpClient), nil
}

// NewClientWithHTTP creates a Hailuo client with a pre-configured HTTP transport.
func NewClientWithHTTP(httpClient core.HTTPClient) *Client {
	return &Client{
		Base:         base.New(httpClient),
		TextToVideo:  &TextToVideo{http: httpClient},
		ImageToVideo: &ImageToVideo{http: httpClient},
	}
}

// TextToVideo generates video from a text prompt.
type TextToVideo struct{ http core.HTTPClient }

// Create submits a Hailuo text-to-video task and returns immediately with a task id.
func (r *TextToVideo) Create(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, textToVideoPath, core.CompactParams(params), requestOptions)
}

// Get fetches the current status of a Hailuo text-to-video task by id.
func (r *TextToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*VideoTaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[VideoTaskResponse](ctx, r.http, core.ResourcePath(textToVideoPath, id), requestOptions)
}

// Run submits a Hailuo text-to-video task and polls until it completes.
func (r *TextToVideo) Run(ctx context.Context, params TextToVideoParams, opts ...option.RequestOption) (*VideoTaskResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*VideoTaskResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}

// ImageToVideo animates a still image into video, guided by a text prompt and first-frame image.
type ImageToVideo struct{ http core.HTTPClient }

// Create submits a Hailuo image-to-video task and returns immediately with a task id.
func (r *ImageToVideo) Create(ctx context.Context, params ImageToVideoParams, opts ...option.RequestOption) (*core.TaskCreateResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.PostJSON[core.TaskCreateResponse](ctx, r.http, imageToVideoPath, core.CompactParams(params), requestOptions)
}

// Get fetches the current status of a Hailuo image-to-video task by id.
func (r *ImageToVideo) Get(ctx context.Context, id string, opts ...option.RequestOption) (*VideoTaskResponse, error) {
	requestOptions, _ := option.ResolveRequestOptions(opts...)
	return core.GetJSON[VideoTaskResponse](ctx, r.http, core.ResourcePath(imageToVideoPath, id), requestOptions)
}

// Run submits a Hailuo image-to-video task and polls until it completes.
func (r *ImageToVideo) Run(ctx context.Context, params ImageToVideoParams, opts ...option.RequestOption) (*VideoTaskResponse, error) {
	_, pollingOptions := option.ResolveRequestOptions(opts...)
	return core.RunAsync(ctx, func(ctx context.Context) (*core.TaskCreateResponse, error) { return r.Create(ctx, params, opts...) }, func(ctx context.Context, id string) (*VideoTaskResponse, error) { return r.Get(ctx, id, opts...) }, pollingOptions)
}
