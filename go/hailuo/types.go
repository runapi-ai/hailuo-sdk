package hailuo

type TextToVideoModel string
type ImageToVideoModel string

const (
	ModelTextToVideoPro      TextToVideoModel  = "hailuo-02-text-to-video-pro"
	ModelTextToVideoStandard TextToVideoModel  = "hailuo-02-text-to-video-standard"
	ModelImageToVideoPro02   ImageToVideoModel = "hailuo-02-image-to-video-pro"
	ModelImageToVideoStd02   ImageToVideoModel = "hailuo-02-image-to-video-standard"
	ModelImageToVideoPro23   ImageToVideoModel = "hailuo-2-3-image-to-video-pro"
	ModelImageToVideoStd23   ImageToVideoModel = "hailuo-2-3-image-to-video-standard"
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
	Model           TextToVideoModel `json:"model" help:"required; hailuo-02-text-to-video-pro or hailuo-02-text-to-video-standard"`
	Prompt          string           `json:"prompt" help:"required; up to 1500 characters"`
	Duration        string           `json:"duration,omitempty" help:"optional; 6 or 10, standard model only"`
	PromptOptimizer *bool            `json:"prompt_optimizer,omitempty" help:"optional; enable Hailuo prompt optimizer"`
	NSFWChecker     *bool            `json:"nsfw_checker,omitempty" help:"optional; content filtering, default false"`
	CallbackURL     string           `json:"callback_url,omitempty" help:"optional; URL that receives completion callback"`
}

type ImageToVideoParams struct {
	Model           ImageToVideoModel `json:"model" help:"required; one of the supported Hailuo image-to-video models"`
	Prompt          string            `json:"prompt" help:"required; 1500 chars on 02 models, 5000 on 2.3 models"`
	ImageURL        string            `json:"image_url" help:"required; publicly accessible image URL"`
	EndImageURL     string            `json:"end_image_url,omitempty" help:"optional; 02 models only"`
	Duration        string            `json:"duration,omitempty" help:"optional; 6 or 10 on standard/2.3 models"`
	Resolution      string            `json:"resolution,omitempty" help:"optional; 512P/768P on 02 standard, 768P/1080P on 2.3"`
	PromptOptimizer *bool             `json:"prompt_optimizer,omitempty" help:"optional; 02 models only"`
	NSFWChecker     *bool             `json:"nsfw_checker,omitempty" help:"optional; content filtering, default false"`
	CallbackURL     string            `json:"callback_url,omitempty" help:"optional; URL that receives completion callback"`
}
