import type { AsyncTaskStatus } from '@runapi.ai/core';

/**
 * Text-to-video model variants. Pro produces higher-fidelity output with longer generation time;
 * Standard is faster with slightly lower quality. Prompts are limited to 1500 characters.
 */
export type HailuoTextToVideoModel =
  | 'hailuo-02-text-to-video-pro'
  | 'hailuo-02-text-to-video-standard';

/**
 * Image-to-video model variants spanning two generations.
 *
 * **02 generation** (`hailuo-02-*`): supports last-frame image control and prompt optimizer.
 * Resolutions: 512p, 768p. Prompts limited to 1500 characters.
 *
 * **2.3 generation** (`hailuo-2.3-*`): higher output resolution (768p, 1080p) and longer prompts
 * (up to 5000 characters), but does not support last-frame image or prompt optimizer.
 *
 * Within each generation, Pro delivers higher fidelity; Standard is faster.
 */
export type HailuoImageToVideoModel =
  | 'hailuo-02-image-to-video-pro'
  | 'hailuo-02-image-to-video-standard'
  | 'hailuo-2.3-image-to-video-pro'
  | 'hailuo-2.3-image-to-video-standard';

/** Video duration in seconds. */
export type HailuoDuration = 6 | 10;
/** Output resolution for 02 generation image-to-video models. */
export type HailuoImage02Resolution = '512p' | '768p';
/** Output resolution for 2.3 generation image-to-video models. 1080p is not available with 10-second duration. */
export type HailuoImage23Resolution = '768p' | '1080p';
/** Combined resolution type across both image-to-video generations. */
export type HailuoImageResolution = HailuoImage02Resolution | HailuoImage23Resolution;

/** Shared parameters for all Hailuo task creation requests. */
interface TaskCommonParams {
  /** Optional webhook URL to receive a notification when the task completes. */
  callback_url?: string;
}

/** Parameters for creating a text-to-video generation task. */
export interface HailuoTextToVideoParams extends TaskCommonParams {
  model: HailuoTextToVideoModel;
  /** Video description prompt (max 1500 characters). */
  prompt: string;
  /** Video length. Defaults to 6 seconds. Not supported on the Pro model. */
  duration_seconds?: HailuoDuration;
  /** Let the model enhance the prompt for better generation results. */
  prompt_optimizer?: boolean;
  /** Run content safety checks on the generated output. */
  enable_safety_checker?: boolean;
}

/**
 * Parameters for creating an image-to-video generation task.
 *
 * Cross-field constraints vary by generation:
 * - **02 Pro**: does not support `duration_seconds` or `output_resolution`.
 * - **02 Standard**: resolution is 512p or 768p; supports `last_frame_image_url` and `prompt_optimizer`.
 * - **2.3 models**: resolution is 768p or 1080p (1080p unavailable at 10 seconds); does not support
 *   `last_frame_image_url` or `prompt_optimizer`. Prompts can be up to 5000 characters.
 */
export interface HailuoImageToVideoParams extends TaskCommonParams {
  model: HailuoImageToVideoModel;
  /** Video motion description prompt. Max 1500 chars for 02 models, 5000 chars for 2.3 models. */
  prompt: string;
  /** Source image URL used as the video's opening frame. */
  first_frame_image_url: string;
  /** Target ending frame image URL. Only supported on 02 generation models. */
  last_frame_image_url?: string;
  /** Video length. Defaults to 6 seconds. Not supported on 02 Pro. */
  duration_seconds?: HailuoDuration;
  /** Output resolution. Allowed values depend on the model generation. */
  output_resolution?: HailuoImageResolution;
  /** Let the model enhance the prompt. Only supported on 02 generation models. */
  prompt_optimizer?: boolean;
  /** Run content safety checks on the generated output. */
  enable_safety_checker?: boolean;
}

/** Response from task creation containing the task identifier for polling. */
export interface TaskCreateResponse {
  id: string;
}

/** A generated video file with a download URL. */
export interface MediaUrl {
  url: string;
}

/** Video generation task status and results. Poll until `status` reaches a terminal state. */
export interface HailuoVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  /** Generated video files, present when the task completes successfully. */
  videos?: MediaUrl[];
  /** Error description if the task failed. */
  error?: string;
  [key: string]: unknown;
}

/** Narrowed response type returned by `run()` after successful completion, guaranteeing `videos` is present. */
export type CompletedHailuoVideoResponse = HailuoVideoResponse & {
  status: 'completed';
  videos: MediaUrl[];
};
