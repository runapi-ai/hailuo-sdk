import { beforeEach, describe, expect, it, vi } from 'vitest';
import type { HttpClient } from '@runapi.ai/core';
import { TextToVideo } from '../../src/resources/text-to-video';
import { ImageToVideo } from '../../src/resources/image-to-video';

describe('Hailuo resources', () => {
  const mockHttp: HttpClient = {
    request: vi.fn(),
  };

  beforeEach(() => {
    vi.clearAllMocks();
  });

  it('creates text-to-video with flat params', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ id: 'task-1' });
    const textToVideo = new TextToVideo(mockHttp);

    await textToVideo.create({
      model: 'hailuo-02-text-to-video-standard',
      prompt: 'A quiet river under moonlight',
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/hailuo/text_to_video', {
      body: {
        model: 'hailuo-02-text-to-video-standard',
        prompt: 'A quiet river under moonlight',
      },
    });
  });

  it('gets text-to-video by id', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({
      id: 'task-1',
      status: 'completed',
      videos: [{ url: 'https://tempfile.runapi.ai/video.mp4' }],
    });
    const textToVideo = new TextToVideo(mockHttp);

    const result = await textToVideo.get('task-1');

    expect(mockHttp.request).toHaveBeenCalledWith('GET', '/api/v1/hailuo/text_to_video/task-1', {});
    expect(result.videos?.[0]?.url).toBe('https://tempfile.runapi.ai/video.mp4');
  });

  it('creates image-to-video with image_url', async () => {
    vi.mocked(mockHttp.request).mockResolvedValueOnce({ id: 'task-2' });
    const imageToVideo = new ImageToVideo(mockHttp);

    await imageToVideo.create({
      model: 'hailuo-2-3-image-to-video-standard',
      prompt: 'Animate the portrait',
      image_url: 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=1200&q=80',
      resolution: '768P',
    });

    expect(mockHttp.request).toHaveBeenCalledWith('POST', '/api/v1/hailuo/image_to_video', {
      body: {
        model: 'hailuo-2-3-image-to-video-standard',
        prompt: 'Animate the portrait',
        image_url: 'https://images.unsplash.com/photo-1506744038136-46273834b3fb?auto=format&fit=crop&w=1200&q=80',
        resolution: '768P',
      },
    });
  });
});
