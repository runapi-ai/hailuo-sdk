package ai.runapi.hailuo.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for image to video operations. */
public final class ImageToVideoParams {
  private final String model;
  private final String prompt;
  private final String firstFrameImageUrl;
  private final String lastFrameImageUrl;
  private final Integer durationSeconds;
  private final String outputResolution;
  private final Boolean promptOptimizer;
  private final Boolean enableSafetyChecker;
  private final String callbackUrl;

  private ImageToVideoParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.firstFrameImageUrl = HailuoParamUtils.requireNonBlank(builder.firstFrameImageUrl, "firstFrameImageUrl");
    this.lastFrameImageUrl = builder.lastFrameImageUrl;
    this.durationSeconds = builder.durationSeconds;
    this.outputResolution = builder.outputResolution;
    this.promptOptimizer = builder.promptOptimizer;
    this.enableSafetyChecker = builder.enableSafetyChecker;
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new ImageToVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "hailuo/image-to-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", HailuoParamUtils.wireValue(model));
    raw.put("prompt", HailuoParamUtils.wireValue(prompt));
    raw.put("first_frame_image_url", HailuoParamUtils.wireValue(firstFrameImageUrl));
    raw.put("last_frame_image_url", HailuoParamUtils.wireValue(lastFrameImageUrl));
    raw.put("duration_seconds", HailuoParamUtils.wireValue(durationSeconds));
    raw.put("output_resolution", HailuoParamUtils.wireValue(outputResolution));
    raw.put("prompt_optimizer", HailuoParamUtils.wireValue(promptOptimizer));
    raw.put("enable_safety_checker", HailuoParamUtils.wireValue(enableSafetyChecker));
    raw.put("callback_url", HailuoParamUtils.wireValue(callbackUrl));
    return HailuoParamUtils.compact(raw);
  }



  /** Builder for {@link ImageToVideoParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private String firstFrameImageUrl;
    private String lastFrameImageUrl;
    private Integer durationSeconds;
    private String outputResolution;
    private Boolean promptOptimizer;
    private Boolean enableSafetyChecker;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(ImageToVideoModel value) {
      this.model = java.util.Objects.requireNonNull(value, "model").value();
      return this;
    }

    /** Sets the model slug using a string value. */
    public Builder model(String value) {
      this.model = HailuoParamUtils.requireNonBlankTrim(value, "model");
      return this;
    }


    /** Sets the text prompt. */
    public Builder prompt(String value) {
      this.prompt = HailuoParamUtils.requireNonBlank(value, "prompt");
      return this;
    }

    /** Sets the first frame image URL. */
    public Builder firstFrameImageUrl(String value) {
      this.firstFrameImageUrl = HailuoParamUtils.requireNonBlank(value, "firstFrameImageUrl");
      return this;
    }

    /** Sets the last frame image URL. */
    public Builder lastFrameImageUrl(String value) {
      this.lastFrameImageUrl = HailuoParamUtils.requireNonBlank(value, "lastFrameImageUrl");
      return this;
    }

    /** Sets the duration in seconds. */
    public Builder durationSeconds(int value) {
      this.durationSeconds = value;
      return this;
    }

    /** Sets the output resolution. */
    public Builder outputResolution(String value) {
      this.outputResolution = HailuoParamUtils.requireNonBlank(value, "outputResolution");
      return this;
    }

    /** Sets the prompt optimizer. */
    public Builder promptOptimizer(boolean value) {
      this.promptOptimizer = value;
      return this;
    }

    /** Sets the content safety checker toggle. */
    public Builder enableSafetyChecker(boolean value) {
      this.enableSafetyChecker = value;
      return this;
    }

    /** Sets the webhook URL for task completion notifications. */
    public Builder callbackUrl(String value) {
      this.callbackUrl = HailuoParamUtils.requireNonBlank(value, "callbackUrl");
      return this;
    }

    /** Builds immutable image to video parameters. */
    public ImageToVideoParams build() {
      return new ImageToVideoParams(this);
    }
  }
}
