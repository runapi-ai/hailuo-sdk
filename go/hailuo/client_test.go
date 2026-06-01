package hailuo

import (
	"context"
	"encoding/json"
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
