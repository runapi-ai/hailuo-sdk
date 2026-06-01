package hailuo

type TextToVideoModel string
type ImageToVideoModel string

const (
	ModelTextToVideoPro      TextToVideoModel  = "hailuo-02-text-to-video-pro"
	ModelTextToVideoStandard TextToVideoModel  = "hailuo-02-text-to-video-standard"
	ModelImageToVideoPro02   ImageToVideoModel = "hailuo-02-image-to-video-pro"
	ModelImageToVideoStd02   ImageToVideoModel = "hailuo-02-image-to-video-standard"
	ModelImageToVideoPro23   ImageToVideoModel = "hailuo-2.3-image-to-video-pro"
	ModelImageToVideoStd23   ImageToVideoModel = "hailuo-2.3-image-to-video-standard"
)

type Video struct {
	URL string `json:"url"`
}

type AsyncTaskResponse struct {
	ID     string `json:"id"`
	Status string `json:"status"`
	Error  string `json:"error,omitempty"`
}

func (r AsyncTaskResponse) GetID() string     { return r.ID }
func (r AsyncTaskResponse) GetStatus() string { return r.Status }
func (r AsyncTaskResponse) GetError() string  { return r.Error }

type VideoTaskResponse struct {
	AsyncTaskResponse
	Videos []Video `json:"videos,omitempty"`
}

type TextToVideoParams struct {
	Model               TextToVideoModel `json:"model" help:"required; model slug"`
	Prompt              string           `json:"prompt" help:"required; up to 1500 characters"`
	DurationSeconds     int              `json:"duration_seconds,omitempty" help:"optional; duration in seconds"`
	PromptOptimizer     *bool            `json:"prompt_optimizer,omitempty" help:"optional; enable Hailuo prompt optimizer"`
	EnableSafetyChecker *bool            `json:"enable_safety_checker,omitempty" help:"optional; content safety check toggle"`
	CallbackURL         string           `json:"callback_url,omitempty" help:"optional; URL that receives completion callback"`
}

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
