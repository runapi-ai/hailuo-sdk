# Hailuo Python SDK for RunAPI

The Hailuo Python SDK is the language-specific package for Hailuo on RunAPI. Use this hailuo ai api package for text-to-video, image-to-video, video editing, and animation flows when your application needs JSON request bodies, task status lookup, and consistent RunAPI errors in Python.

This hailuo ai api README is the Python package guide inside the public `hailuo-sdk` repository. For the repository overview, start at `../README.md`; for model details, use https://runapi.ai/models/hailuo; for API reference, use https://runapi.ai/docs#hailuo; for SDK docs, use https://runapi.ai/docs#sdk-hailuo.

## Install

```bash
pip install runapi-hailuo
```

## Quick start

```python
from runapi.hailuo import HailuoClient

client = HailuoClient()  # reads RUNAPI_API_KEY, or pass api_key="sk-..."

task = client.text_to_video.create(
    model="hailuo-02-text-to-video-standard",
    prompt="A cinematic drone shot over a coastal city at dusk",
)
status = client.text_to_video.get(task.id)

video = client.image_to_video.create(
    model="hailuo-2.3-image-to-video-standard",
    prompt="Slowly zoom in as the waves roll",
    first_frame_image_url="https://example.com/source.jpg",
)
```

Use `create` to submit a task and return quickly, `get` to fetch the latest task state, and `run` to create and poll until completion:

```python
result = client.text_to_video.run(
    model="hailuo-02-text-to-video-standard",
    prompt="A futuristic cityscape at dusk",
)
print(result.videos[0].url)
```

In web request handlers, prefer `create` plus webhook or later `get` polling so a worker is not held open.

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## Language notes

Pass parameters as keyword arguments and catch the `runapi.hailuo` error classes when building video jobs or scripts. The available resources are `text_to_video` and `image_to_video`. Keep `RUNAPI_API_KEY` in the environment or your secret manager; never commit API keys or callback secrets.

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
