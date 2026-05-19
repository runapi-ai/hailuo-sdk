import { describe, expect, it } from 'vitest';
import { HailuoClient } from '../src';

describe('HailuoClient', () => {
  it('exposes the two resources', () => {
    const client = new HailuoClient({ apiKey: 'test-key' });

    expect(client.textToVideo).toBeDefined();
    expect(client.imageToVideo).toBeDefined();
  });
});
