<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/hailuo-sdk">Hailuo API SDK for RunAPI</a>
</h3>

<p align="center">
  Hailuo API SDKs for JavaScript, Ruby, and Go on RunAPI.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/hailuo)](https://www.npmjs.com/package/@runapi.ai/hailuo)
[![RubyGems](https://img.shields.io/gem/v/runapi-hailuo)](https://rubygems.org/gems/runapi-hailuo)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/hailuo-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/hailuo-sdk/go)
[![License](https://img.shields.io/github/license/runapi-ai/hailuo-sdk)](https://github.com/runapi-ai/hailuo-sdk/blob/main/LICENSE)

</div>
<br/>

The hailuo ai api SDK packages JavaScript, Ruby, and Go clients for Hailuo on RunAPI. Use this hailuo ai api SDK for text-to-video and image-to-video workflows that need typed installs, JSON request bodies, task polling, and consistent RunAPI errors across services.

Hailuo belongs to the MiniMax catalog on RunAPI. The public model page is https://runapi.ai/models/hailuo; variant pages below carry pricing, rate-limit, and commercial-usage details. The public `hailuo-sdk` repository groups the JavaScript, Ruby, and Go packages for this model.

## Install

```bash
npm install @runapi.ai/hailuo
gem install runapi-hailuo
go get github.com/runapi-ai/hailuo-sdk/go@latest
```

## What you can build

- Build marketing clips, storyboard previews, creator tools, and agent video pipelines with the hailuo ai api SDK.
- Keep one model-specific repository while installing only the language package your app needs.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Handle authentication, validation, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

The JavaScript client exposes text to videos, image to videos resources, and the Ruby and Go packages mirror the same RunAPI task lifecycle.

## JavaScript quick start

```typescript
import { HailuoClient } from '@runapi.ai/hailuo';

const client = new HailuoClient();

const task = await client.textToVideo.create({
  // Pass the Hailuo request body documented at https://runapi.ai/docs#hailuo.
});

const status = await client.textToVideo.get(task.id);
```

For short scripts, use `run` with the same JSON body to create the task and wait for completion. For web request handlers, prefer `create` plus webhook or later `get` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/hailuo`.
- `ruby/` publishes `runapi-hailuo` when RubyGems publishing resumes.
- `go/` publishes `github.com/runapi-ai/hailuo-sdk/go` and depends on `github.com/runapi-ai/core-sdk/go`.

## Public links

- Model page: https://runapi.ai/models/hailuo
- SDK docs: https://runapi.ai/docs#sdk-hailuo
- Product docs: https://runapi.ai/docs#hailuo
- SDK repository: https://github.com/runapi-ai/hailuo-sdk
- Skill repository: https://github.com/runapi-ai/hailuo
- Provider comparison: https://runapi.ai/providers/minimax
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific hailuo ai api variant page for pricing, rate limits, and commercial usage:
- [02 text to video pro](https://runapi.ai/models/hailuo/02-text-to-video-pro)
- [02 text to video standard](https://runapi.ai/models/hailuo/02-text-to-video-standard)
- [02 image to video pro](https://runapi.ai/models/hailuo/02-image-to-video-pro)
- [02 image to video standard](https://runapi.ai/models/hailuo/02-image-to-video-standard)
- [2.3 image to video pro](https://runapi.ai/models/hailuo/2.3-image-to-video-pro)
- [2.3 image to video standard](https://runapi.ai/models/hailuo/2.3-image-to-video-standard)

Default pricing link for the hailuo ai api SDK: https://runapi.ai/models/hailuo/02-text-to-video-pro

## Generated file storage

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## FAQ

### Which package should I install for hailuo ai api work?

Install the model package for your language: `@runapi.ai/hailuo`, `runapi-hailuo`, or `github.com/runapi-ai/hailuo-sdk/go`. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary hailuo ai api links point to https://runapi.ai/models/hailuo. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/hailuo/02-text-to-video-pro. Provider comparisons point to https://runapi.ai/providers/minimax, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
