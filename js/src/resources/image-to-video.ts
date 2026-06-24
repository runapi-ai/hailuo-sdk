import type { HttpClient, PollingOptions, RequestOptions, ActionSchema } from '@runapi.ai/core';
import { compactParams, validateParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import { contract } from '../contract_gen';
import type {
  CompletedHailuoVideoResponse,
  HailuoImageToVideoParams,
  HailuoVideoResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/hailuo/image_to_video';

/** Animate a still image into video, guided by a text prompt and first-frame image. Supports two model generations with different resolution and feature sets. */
export class ImageToVideo {
  constructor(private readonly http: HttpClient) {}

  /**
   * Generate a video and wait until complete.
   * @param params Image-to-video parameters.
   * @param options Per-request and polling overrides.
   * @returns The completed task with videos.
   */
  async run(
    params: HailuoImageToVideoParams,
    options?: RequestOptions & PollingOptions
  ): Promise<CompletedHailuoVideoResponse> {
    const { id } = await this.create(params, options);
    const response = await pollUntilComplete<HailuoVideoResponse>(() => this.get(id, options), {
      maxWaitMs: options?.maxWaitMs,
      pollIntervalMs: options?.pollIntervalMs,
    });
    return response as CompletedHailuoVideoResponse;
  }

  /**
   * Create an image-to-video task; returns immediately with a task id.
   * @param params Image-to-video parameters.
   * @param options Per-request overrides.
   * @returns The task creation result with id.
   */
  async create(params: HailuoImageToVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    const body = compactParams(params);
    validateParams(contract['image-to-video'] as ActionSchema, body as Record<string, unknown>);
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body,
      ...options,
    });
  }

  /**
   * Fetch the current status of an image-to-video task.
   * @param id The task id.
   * @param options Per-request overrides.
   * @returns The current task status.
   */
  async get(id: string, options?: RequestOptions): Promise<HailuoVideoResponse> {
    return this.http.request<HailuoVideoResponse>('GET', `${ENDPOINT}/${id}`, options ?? {});
  }
}
