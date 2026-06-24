package ai.runapi.hailuo.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for text to video operations. */
public final class TextToVideoModel extends HailuoValue {
  /** hailuo-02-text-to-video-pro model slug. */
  public static final TextToVideoModel HAILUO_02_TEXT_TO_VIDEO_PRO = new TextToVideoModel("hailuo-02-text-to-video-pro");
  /** hailuo-02-text-to-video-standard model slug. */
  public static final TextToVideoModel HAILUO_02_TEXT_TO_VIDEO_STANDARD = new TextToVideoModel("hailuo-02-text-to-video-standard");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public TextToVideoModel(String value) {
    super(value);
  }
}
