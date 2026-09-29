package hailuo

import (
	"context"
	"encoding/json"
	"strings"
	"testing"

	"github.com/runapi-ai/core-sdk/go/core"
)

type stubHTTPClient struct {
	method string
	path   string
	body   any
}

func (s *stubHTTPClient) Request(_ context.Context, method, path string, opts *core.HTTPRequestOptions) (json.RawMessage, error) {
	s.method = method
	s.path = path
	if opts != nil {
		s.body = opts.Body
	}
	return json.RawMessage(`{"id":"task_123","status":"processing"}`), nil
}

func TestTextToVideoCreate(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	_, err := client.TextToVideo.Create(context.Background(), TextToVideoParams{
		Model:  ModelTextToVideoStandard,
		Prompt: "A quiet river under moonlight",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/hailuo/text_to_video" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["model"] != "hailuo-02-text-to-video-standard" {
		t.Fatalf("unexpected model: %v", body["model"])
	}
}

func TestImageToVideoRejects23TenSecond1080p(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	_, err := client.ImageToVideo.Create(context.Background(), ImageToVideoParams{
		Model:              ModelImageToVideoStd23,
		Prompt:             "Animate the portrait",
		FirstFrameImageURL: "https://cdn.runapi.ai/public/samples/input.png",
		DurationSeconds:    10,
		OutputResolution:   "1080p",
	})
	want := "output_resolution must be one of: 768p when duration_seconds is 10 and model is hailuo-2.3-image-to-video-standard"
	if err == nil || !strings.Contains(err.Error(), want) {
		t.Fatalf("unexpected error: %v", err)
	}
	if stub.body != nil {
		t.Fatalf("expected no request body, got: %v", stub.body)
	}
}

func TestImageToVideoCreate(t *testing.T) {
	stub := &stubHTTPClient{}
	client := NewClientWithHTTP(stub)
	_, err := client.ImageToVideo.Create(context.Background(), ImageToVideoParams{
		Model:              ModelImageToVideoStd23,
		Prompt:             "Animate the portrait",
		FirstFrameImageURL: "https://cdn.runapi.ai/public/samples/input.png",
		OutputResolution:   "768p",
	})
	if err != nil {
		t.Fatal(err)
	}
	if stub.method != "POST" || stub.path != "/api/v1/hailuo/image_to_video" {
		t.Fatalf("unexpected request: %s %s", stub.method, stub.path)
	}
	body := stub.body.(map[string]any)
	if body["first_frame_image_url"] != "https://cdn.runapi.ai/public/samples/input.png" {
		t.Fatalf("unexpected first_frame_image_url: %v", body["first_frame_image_url"])
	}
	if body["output_resolution"] != "768p" {
		t.Fatalf("unexpected output_resolution: %v", body["output_resolution"])
	}
	if _, ok := body["resolution"]; ok {
		t.Fatalf("unexpected resolution key in body: %#v", body)
	}
	if _, ok := body["image_url"]; ok {
		t.Fatalf("unexpected image_url key in body: %#v", body)
	}
}
