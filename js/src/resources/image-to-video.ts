import type { HttpClient, PollingOptions, RequestOptions } from '@runapi.ai/core';
import { compactParams } from '@runapi.ai/core';
import { pollUntilComplete } from '@runapi.ai/core/internal';
import type {
  CompletedHailuoVideoResponse,
  HailuoImageToVideoParams,
  HailuoVideoResponse,
  TaskCreateResponse,
} from '../types';

const ENDPOINT = '/api/v1/hailuo/image_to_video';

export class ImageToVideo {
  constructor(private readonly http: HttpClient) {}

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

  async create(params: HailuoImageToVideoParams, options?: RequestOptions): Promise<TaskCreateResponse> {
    return this.http.request<TaskCreateResponse>('POST', ENDPOINT, {
      body: compactParams(params),
      ...options,
    });
  }

  async get(id: string, options?: RequestOptions): Promise<HailuoVideoResponse> {
    return this.http.request<HailuoVideoResponse>('GET', `${ENDPOINT}/${id}`, options ?? {});
  }
}
