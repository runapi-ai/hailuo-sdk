package ai.runapi.hailuo.types;

import com.fasterxml.jackson.annotation.JsonCreator;

/** Model slug for image to video operations. */
public final class ImageToVideoModel extends HailuoValue {
  /** hailuo-02-image-to-video-pro model slug. */
  public static final ImageToVideoModel HAILUO_02_IMAGE_TO_VIDEO_PRO = new ImageToVideoModel("hailuo-02-image-to-video-pro");
  /** hailuo-02-image-to-video-standard model slug. */
  public static final ImageToVideoModel HAILUO_02_IMAGE_TO_VIDEO_STANDARD = new ImageToVideoModel("hailuo-02-image-to-video-standard");
  /** hailuo-2.3-image-to-video-pro model slug. */
  public static final ImageToVideoModel HAILUO_2_3_IMAGE_TO_VIDEO_PRO = new ImageToVideoModel("hailuo-2.3-image-to-video-pro");
  /** hailuo-2.3-image-to-video-standard model slug. */
  public static final ImageToVideoModel HAILUO_2_3_IMAGE_TO_VIDEO_STANDARD = new ImageToVideoModel("hailuo-2.3-image-to-video-standard");

  /** Creates a model value from a literal model slug. */
  @JsonCreator
  public ImageToVideoModel(String value) {
    super(value);
  }
}
