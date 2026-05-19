# Hailuo AI API Skill for RunAPI

Generate video with Hailuo 02 and 2.3 text-to-video and image-to-video in Pro and Standard modes. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate Hailuo through RunAPI.

The canonical agent file is `skills/hailuo/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/hailuo -g
```

Or manually: clone this repo and copy `skills/hailuo/` into your agent's skills directory.

## Quick example

```typescript
import { HailuoClient } from '@runapi.ai/hailuo';

const client = new HailuoClient();
const result = await client.textToVideo.run({
  model: 'hailuo-02-text-to-video-pro',
  prompt: 'A cinematic sunset over the ocean',
});
```

## Routing

- Model page: https://runapi.ai/models/hailuo
- Product docs: https://runapi.ai/docs#hailuo
- SDK docs: https://runapi.ai/docs#sdk-hailuo
- SDK repository: https://github.com/runapi-ai/hailuo-sdk
- Pricing and rate limits: https://runapi.ai/models/hailuo/02-text-to-video-pro
- Provider comparison: https://runapi.ai/providers/minimax
- Browse all RunAPI models and skills: https://runapi.ai/models

## Variants

- [02 text to video pro](https://runapi.ai/models/hailuo/02-text-to-video-pro)
- [02 text to video standard](https://runapi.ai/models/hailuo/02-text-to-video-standard)
- [02 image to video pro](https://runapi.ai/models/hailuo/02-image-to-video-pro)
- [02 image to video standard](https://runapi.ai/models/hailuo/02-image-to-video-standard)
- [2.3 image to video pro](https://runapi.ai/models/hailuo/2.3-image-to-video-pro)
- [2.3 image to video standard](https://runapi.ai/models/hailuo/2.3-image-to-video-standard)

## Agent rules

- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For hailuo ai api pricing, rate-limit, and commercial-usage answers, link to the variant page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
