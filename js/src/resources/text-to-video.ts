import type { HttpClient, PollingOptions, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type {
  CompletedHailuoVideoResponse,
  HailuoTextToVideoParams,
  HailuoVideoResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/hailuo/text_to_video';

/** Generate video from a text prompt. Pro delivers higher fidelity; Standard is faster. */
export class TextToVideo {
  constructor(private readonly http: HttpClient) {}

  /**
   * Generate a video and wait until complete.
   * @param params Text-to-video parameters.
   * @param options Per-request and polling overrides.
   * @returns The completed task with videos.
   */
  async run(
    params: HailuoTextToVideoParams,
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
   * Create a text-to-video task; returns immediately with a task id.
   * @param params Text-to-video parameters.
   * @param options Per-request overrides.
   * @returns The task creation result with id.
   */
  async create(params: HailuoTextToVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  /**
   * Fetch the current status of a text-to-video task.
   * @param id The task id.
   * @param options Per-request overrides.
   * @returns The current task status.
   */
  async get(id: string, options?: RequestOptions): Promise<HailuoVideoResponse> {
    return this.http.request<HailuoVideoResponse>('GET', `${ENDPOINT}/${id}`, options ?? {});
  }
}
