<p align="center">
  <a href="https://github.com/runapi-ai/hailuo">
    <h3 align="center">Hailuo AI API Skill for RunAPI</h3>
  </a>
</p>

<p align="center">
  Install this agent skill, inspect Hailuo fields, then run jobs through the RunAPI CLI.
</p>

<p align="center">
  <a href="https://runapi.ai/models/hailuo"><strong>Model Reference</strong></a> · <a href="https://github.com/runapi-ai/cli"><strong>CLI</strong></a> · <a href="https://github.com/runapi-ai/hailuo-sdk"><strong>SDK</strong></a>
</p>

<div align="center">

[![skills.sh](https://www.skills.sh/b/runapi-ai/hailuo)](https://www.skills.sh/runapi-ai/hailuo/hailuo)
[![ClawHub](https://img.shields.io/badge/ClawHub-runapi--hailuo-111827)](https://clawhub.ai/runapi-ai/runapi-hailuo)
[![License](https://img.shields.io/github/license/runapi-ai/hailuo)](https://github.com/runapi-ai/hailuo/blob/main/LICENSE)

</div>
<br/>

Generate video with Hailuo 02 and 2.3 text-to-video and image-to-video in Pro and Standard modes. This skill helps Claude Code, Codex, Gemini CLI, Cursor, and 50+ agents integrate Hailuo through RunAPI.

The canonical agent file is `skills/hailuo/SKILL.md`.

## Install

```bash
npx skills add runapi-ai/hailuo -g
```

Or paste this prompt to your AI agent:

```text
Install the hailuo skill for me:

1. Clone https://github.com/runapi-ai/hailuo
2. Copy the skills/hailuo/ directory into your
   user-level skills directory (e.g. ~/.claude/skills/
   for Claude Code, ~/.codex/skills/ for Codex).
3. Verify that SKILL.md is present.
4. Confirm the install path when done.
```

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

- Integration work uses the target language SDK; one-off generation, manual smoke tests, debugging, or user-requested CLI runs use the RunAPI CLI skill: https://github.com/runapi-ai/cli-skill
- RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.
- Keep API keys in `RUNAPI_API_KEY` or RunAPI CLI config; never commit secrets.
- Prefer `create`, `get`, and `run` JSON passthrough patterns instead of inventing flags for every model parameter.
- For hailuo ai api pricing, rate-limit, and commercial-usage answers, link to the variant page rather than the repository README.

## License

Licensed under the Apache License, Version 2.0.
