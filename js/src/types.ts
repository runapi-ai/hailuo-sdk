import type { AsyncTaskStatus } from '@runapi.ai/core';

export type HailuoTextToVideoModel =
  | 'hailuo-02-text-to-video-pro'
  | 'hailuo-02-text-to-video-standard';

export type HailuoImageToVideoModel =
  | 'hailuo-02-image-to-video-pro'
  | 'hailuo-02-image-to-video-standard'
  | 'hailuo-2.3-image-to-video-pro'
  | 'hailuo-2.3-image-to-video-standard';

export type HailuoDuration = '6' | '10';
export type HailuoImage02Resolution = '512P' | '768P';
export type HailuoImage23Resolution = '768P' | '1080P';
export type HailuoImageResolution = HailuoImage02Resolution | HailuoImage23Resolution;

interface TaskCommonParams {
  callback_url?: string;
}

export interface HailuoTextToVideoParams extends TaskCommonParams {
  model: HailuoTextToVideoModel;
  prompt: string;
  duration?: HailuoDuration | 6 | 10;
  prompt_optimizer?: boolean;
  nsfw_checker?: boolean;
}

export interface HailuoImageToVideoParams extends TaskCommonParams {
  model: HailuoImageToVideoModel;
  prompt: string;
  image_url: string;
  end_image_url?: string;
  duration?: HailuoDuration | 6 | 10;
  resolution?: HailuoImageResolution;
  prompt_optimizer?: boolean;
  nsfw_checker?: boolean;
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
