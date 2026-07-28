package hailuo

import "github.com/runapi-ai/core-sdk/go/core"

// TextToVideoModel selects the Hailuo text-to-video engine variant.
type TextToVideoModel string

// ImageToVideoModel selects the Hailuo image-to-video engine variant.
// The 02 and 2.3 generations differ in prompt length limits and supported features.
type ImageToVideoModel string

const (
	// ModelTextToVideoPro produces higher-fidelity video at the cost of longer generation time.
	ModelTextToVideoPro TextToVideoModel = "hailuo-02-text-to-video-pro"
	// ModelTextToVideoStandard is faster with slightly lower visual quality than Pro.
	ModelTextToVideoStandard TextToVideoModel = "hailuo-02-text-to-video-standard"
	// ModelImageToVideoPro02 is the 02-generation Pro model. Supports last-frame image control and 1500-char prompts.
	ModelImageToVideoPro02 ImageToVideoModel = "hailuo-02-image-to-video-pro"
	// ModelImageToVideoStd02 is the 02-generation Standard model. Same feature set as Pro02 with faster generation.
	ModelImageToVideoStd02 ImageToVideoModel = "hailuo-02-image-to-video-standard"
	// ModelImageToVideoPro23 is the 2.3-generation Pro model. Supports 5000-char prompts and output resolution control, but not last-frame images.
	ModelImageToVideoPro23 ImageToVideoModel = "hailuo-2.3-image-to-video-pro"
	// ModelImageToVideoStd23 is the 2.3-generation Standard model. Same feature set as Pro23 with faster generation.
	ModelImageToVideoStd23 ImageToVideoModel = "hailuo-2.3-image-to-video-standard"
)

// Video holds a URL to a generated video file.
type Video struct {
	URL string `json:"url"`
}

// AsyncTaskResponse carries the task ID, lifecycle status, and error for all Hailuo async operations.
type AsyncTaskResponse struct {
	core.TaskBillingFacts
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return r.Status }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

// VideoTaskResponse is the completed result of a text-to-video or image-to-video task.
type VideoTaskResponse struct {
	AsyncTaskResponse
	Videos []Video `json:"videos,omitempty"`
}

// TextToVideoParams configures text-to-video generation.
// Prompts are limited to 1500 characters.
type TextToVideoParams struct {
	Model               TextToVideoModel `json:"model" help:"required; model slug"`
	Prompt              string           `json:"prompt" help:"required; up to 1500 characters"`
	DurationSeconds     int              `json:"duration_seconds,omitempty" help:"optional; duration in seconds"`
	PromptOptimizer     *bool            `json:"prompt_optimizer,omitempty" help:"optional; enable Hailuo prompt optimizer"`
	EnableSafetyChecker *bool            `json:"enable_safety_checker,omitempty" help:"optional; content safety check toggle"`
	CallbackURL         string           `json:"callback_url,omitempty" help:"optional; URL that receives completion callback"`
}

// ImageToVideoParams configures image-to-video generation.
// Feature availability varies by generation: 02 models support LastFrameImageURL and PromptOptimizer;
// 2.3 models support OutputResolution and 5000-char prompts (vs 1500 for 02).
type ImageToVideoParams struct {
	Model               ImageToVideoModel `json:"model" help:"required; model slug"`
	Prompt              string            `json:"prompt" help:"required; 1500 chars on 02 models, 5000 on 2.3 models"`
	FirstFrameImageURL  string            `json:"first_frame_image_url" help:"required; first-frame image URL"`
	LastFrameImageURL   string            `json:"last_frame_image_url,omitempty" help:"optional; last-frame image URL, 02 models only"`
	DurationSeconds     int               `json:"duration_seconds,omitempty" help:"optional; duration in seconds"`
	OutputResolution    string            `json:"output_resolution,omitempty" help:"optional; output resolution"`
	PromptOptimizer     *bool             `json:"prompt_optimizer,omitempty" help:"optional; 02 models only"`
	EnableSafetyChecker *bool             `json:"enable_safety_checker,omitempty" help:"optional; content safety check toggle"`
	CallbackURL         string            `json:"callback_url,omitempty" help:"optional; URL that receives completion callback"`
}
