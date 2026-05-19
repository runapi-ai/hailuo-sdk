import { createHttpClient, type ClientOptions } from '@runapi.ai/core';
import { TextToVideo } from './resources/text-to-video';
import { ImageToVideo } from './resources/image-to-video';

export class HailuoClient {
  public readonly textToVideo: TextToVideo;
  public readonly imageToVideo: ImageToVideo;

  constructor(options: ClientOptions = {}) {
    const http = createHttpClient(options);
    this.textToVideo = new TextToVideo(http);
    this.imageToVideo = new ImageToVideo(http);
  }
}
