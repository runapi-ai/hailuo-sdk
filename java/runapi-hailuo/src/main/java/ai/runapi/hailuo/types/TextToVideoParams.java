package ai.runapi.hailuo.types;

import java.util.LinkedHashMap;
import java.util.List;
import java.util.Map;

/** Parameters for text to video operations. */
public final class TextToVideoParams {
  private final String model;
  private final String prompt;
  private final Integer durationSeconds;
  private final Boolean promptOptimizer;
  private final Boolean enableSafetyChecker;
  private final String callbackUrl;

  private TextToVideoParams(Builder builder) {
    this.model = builder.model;
    this.prompt = builder.prompt;
    this.durationSeconds = builder.durationSeconds;
    this.promptOptimizer = builder.promptOptimizer;
    this.enableSafetyChecker = builder.enableSafetyChecker;
    this.callbackUrl = builder.callbackUrl;
  }

  /** Creates a new TextToVideoParams builder. */
  public static Builder builder() {
    return new Builder();
  }

  /** Returns the RunAPI action key for this request. */
  public String action() {
    return "hailuo/text-to-video";
  }

  /** Converts these parameters to the JSON request body shape. */
  public Map<String, Object> toMap() {
    Map<String, Object> raw = new LinkedHashMap<String, Object>();
    raw.put("model", HailuoParamUtils.wireValue(model));
    raw.put("prompt", HailuoParamUtils.wireValue(prompt));
    raw.put("duration_seconds", HailuoParamUtils.wireValue(durationSeconds));
    raw.put("prompt_optimizer", HailuoParamUtils.wireValue(promptOptimizer));
    raw.put("enable_safety_checker", HailuoParamUtils.wireValue(enableSafetyChecker));
    raw.put("callback_url", HailuoParamUtils.wireValue(callbackUrl));
    return HailuoParamUtils.compact(raw);
  }



  /** Builder for {@link TextToVideoParams}. */
  public static final class Builder {
    private String model;
    private String prompt;
    private Integer durationSeconds;
    private Boolean promptOptimizer;
    private Boolean enableSafetyChecker;
    private String callbackUrl;

    private Builder() {}

    /** Sets the model slug using a typed model value. */
    public Builder model(TextToVideoModel value) {
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

    /** Sets the duration in seconds. */
    public Builder durationSeconds(int value) {
      this.durationSeconds = value;
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

    /** Builds immutable text to video parameters. */
    public TextToVideoParams build() {
      return new TextToVideoParams(this);
    }
  }
}
