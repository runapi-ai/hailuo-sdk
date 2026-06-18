# Hailuo AI API Ruby SDK for RunAPI

The hailuo ai api Ruby SDK is the language-specific package for Hailuo on RunAPI. Use this hailuo ai api package for text-to-video, image-to-video, video editing, and animation flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Ruby.

This hailuo ai api README is the Ruby package guide inside the public `hailuo-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/hailuo; for API reference, use https://runapi.ai/docs#hailuo; for SDK docs, use https://runapi.ai/docs#sdk-hailuo.

## Install

```bash
gem install runapi-hailuo
```

## Quick start

```ruby
require "runapi-hailuo"

client = RunApi::Hailuo::Client.new
task = client.text_to_video.create(
  # Pass the Hailuo JSON request body from https://runapi.ai/docs#hailuo.
)
status = client.text_to_video.get(task.id)
```

Use `create` when you want to submit a task and return quickly, `get` when you need the latest task state, and `run` when a script should create and poll until completion. In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Use Ruby keyword arguments and the `RunApi::Hailuo` error classes when building video jobs, Rails workers, or scripts. The available resources include text to videos, and image to videos. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
