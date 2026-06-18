import { BaseClient, type ClientOptions } from '@runapi.ai/core';
import { TextToVideo } from './resources/text-to-video';
import { ImageToVideo } from './resources/image-to-video';

/**
 * Hailuo text-to-video and image-to-video generation API client.
 *
 * @example
 * ```typescript
 * const client = new HailuoClient({ apiKey: 'your-api-key' });
 *
 * const result = await client.textToVideo.run({
 *   model: 'hailuo-02-text-to-video-standard',
 *   prompt: 'A timelapse of cherry blossoms blooming in a Japanese garden',
 * });
 * ```
 */
export class HailuoClient extends BaseClient {
  /** Text-to-video generation from a text prompt. Pro delivers higher fidelity; Standard is faster. */
  public readonly textToVideo: TextToVideo;
  /** Image-to-video animation from a first-frame image guided by a text prompt. */
  public readonly imageToVideo: ImageToVideo;

  constructor(options: ClientOptions = {}) {
    super(options);
    this.textToVideo = new TextToVideo(this.http);
    this.imageToVideo = new ImageToVideo(this.http);
  }
}
