# Hailuo AI API JavaScript SDK for RunAPI

The hailuo ai api JavaScript SDK is the language-specific package for Hailuo on RunAPI. Use this hailuo ai api package for text-to-video, image-to-video, video-to-video, animation, and edit flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in JavaScript.

This hailuo ai api README is the JavaScript package guide inside the public `hailuo-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/hailuo; for API reference, use https://runapi.ai/docs#hailuo; for SDK docs, use https://runapi.ai/docs#sdk-hailuo.

## Install

```bash
npm install @runapi.ai/hailuo
```

## Quick start

```typescript
import { HailuoClient } from '@runapi.ai/hailuo';

const client = new HailuoClient();
const task = await client.textToVideo.create({
  // Pass the Hailuo JSON request body from https://runapi.ai/docs#hailuo.
});
const status = await client.textToVideo.get(task.id);
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

## Language notes

Use the TypeScript types in `src/types.ts` and the resource classes under `src/resources` when building video applications. The available resources include text to videos, and image to videos. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

## Links

- Model page: https://runapi.ai/models/hailuo
- SDK docs: https://runapi.ai/docs#sdk-hailuo
- Product docs: https://runapi.ai/docs#hailuo
- Pricing and rate limits: https://runapi.ai/models/hailuo/02-text-to-video-pro
- Provider comparison: https://runapi.ai/providers/minimax
- Full catalog: https://runapi.ai/models
- Repository: https://github.com/runapi-ai/hailuo-sdk

## License

Licensed under the Apache License, Version 2.0.
