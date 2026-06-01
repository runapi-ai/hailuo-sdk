import type { AsyncTaskStatus } from '@runapi.ai/core';

export type HailuoTextToVideoModel =
  | 'hailuo-02-text-to-video-pro'
  | 'hailuo-02-text-to-video-standard';

export type HailuoImageToVideoModel =
  | 'hailuo-02-image-to-video-pro'
  | 'hailuo-02-image-to-video-standard'
  | 'hailuo-2.3-image-to-video-pro'
  | 'hailuo-2.3-image-to-video-standard';

export type HailuoDuration = 6 | 10;
export type HailuoImage02Resolution = '512p' | '768p';
export type HailuoImage23Resolution = '768p' | '1080p';
export type HailuoImageResolution = HailuoImage02Resolution | HailuoImage23Resolution;

interface TaskCommonParams {
  callback_url?: string;
}

export interface HailuoTextToVideoParams extends TaskCommonParams {
  model: HailuoTextToVideoModel;
  prompt: string;
  duration_seconds?: HailuoDuration;
  prompt_optimizer?: boolean;
  enable_safety_checker?: boolean;
}

export interface HailuoImageToVideoParams extends TaskCommonParams {
  model: HailuoImageToVideoModel;
  prompt: string;
  first_frame_image_url: string;
  last_frame_image_url?: string;
  duration_seconds?: HailuoDuration;
  output_resolution?: HailuoImageResolution;
  prompt_optimizer?: boolean;
  enable_safety_checker?: boolean;
}

export interface TaskCreateResponse {
  id: string;
}

export interface MediaUrl {
  url: string;
}

export interface HailuoVideoResponse {
  id: string;
  status: AsyncTaskStatus;
  videos?: MediaUrl[];
  error?: string;
  [key: string]: unknown;
}

export type CompletedHailuoVideoResponse = HailuoVideoResponse & {
  status: 'completed';
  videos: MediaUrl[];
};
