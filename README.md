<p align="center">
  <a href="https://runapi.ai"><img src="https://runapi.ai/icon.svg" height="56" alt="RunAPI"></a>
</p>

<h3 align="center">
  <a href="https://github.com/runapi-ai/hailuo-sdk">Hailuo API SDK for RunAPI</a>
</h3>

<p align="center">
  Hailuo API SDKs for JavaScript, Python, Ruby, Go, Java, and PHP on RunAPI.
</p>

<div align="center">

[![npm](https://img.shields.io/npm/v/@runapi.ai/hailuo)](https://www.npmjs.com/package/@runapi.ai/hailuo)
[![PyPI](https://img.shields.io/pypi/v/runapi-hailuo)](https://pypi.org/project/runapi-hailuo/)
[![RubyGems](https://img.shields.io/gem/v/runapi-hailuo)](https://rubygems.org/gems/runapi-hailuo)
[![Go Reference](https://pkg.go.dev/badge/github.com/runapi-ai/hailuo-sdk/go.svg)](https://pkg.go.dev/github.com/runapi-ai/hailuo-sdk/go)
[![Maven Central](https://img.shields.io/maven-central/v/ai.runapi/runapi-hailuo)](https://central.sonatype.com/artifact/ai.runapi/runapi-hailuo)
[![License](https://img.shields.io/github/license/runapi-ai/hailuo-sdk)](https://github.com/runapi-ai/hailuo-sdk/blob/main/LICENSE)

</div>
<br/>

The Hailuo API SDK packages JavaScript, Python, Ruby, Go, Java, and PHP clients for Hailuo on RunAPI. Use it for text-to-video and image-to-video workflows when your app needs typed request builders, predictable task polling, file upload helpers, account helpers, and consistent RunAPI errors.

Hailuo is listed in the RunAPI model catalog at https://runapi.ai/models/hailuo. Variant pages below carry pricing, rate-limit, and commercial-usage details. The public `hailuo-sdk` repository groups the non-PHP language packages, examples, CI, and release tags for this model. The PHP package is released from a split Composer repository.

## Install

```bash
npm install @runapi.ai/hailuo
pip install runapi-hailuo
gem install runapi-hailuo
go get github.com/runapi-ai/hailuo-sdk/go@latest
```

Gradle:

```kotlin
dependencies {
  implementation("ai.runapi:runapi-hailuo:0.1.1")
}
```

Maven:

```xml
<dependency>
  <groupId>ai.runapi</groupId>
  <artifactId>runapi-hailuo</artifactId>
  <version>0.1.1</version>
</dependency>
```

Use the Java BOM when installing multiple RunAPI Java modules:

```kotlin
dependencies {
  implementation(platform("ai.runapi:runapi-bom:0.2.7"))
  implementation("ai.runapi:runapi-hailuo")
}
```

The PHP package is published from the split Composer repository as `runapi-ai/hailuo`; see https://github.com/runapi-ai/hailuo-php for PHP install and examples.

## What you can build

- Build apps, agent workflows, batch jobs, and production services around Hailuo requests.
- Install only the language package your app needs while keeping one model-specific repository for docs and releases.
- Use `create` for submit-only jobs, `get` for status lookup, and `run` for submit-and-poll scripts.
- Upload local files, URL files, or base64 files through shared RunAPI file helpers.
- Handle validation, authentication, rate limits, insufficient credits, task failures, and polling timeouts through RunAPI SDK errors.

## Java quick start

```java
import ai.runapi.hailuo.HailuoClient;
import ai.runapi.hailuo.types.TextToVideoParams;
import ai.runapi.hailuo.types.CompletedTextToVideoResponse;
import ai.runapi.hailuo.types.TextToVideoModel;

HailuoClient client = HailuoClient.builder()
    .apiKey(System.getenv("RUNAPI_API_KEY"))
    .build();

CompletedTextToVideoResponse result = client.textToVideo().run(
    TextToVideoParams.builder()
        .model(TextToVideoModel.HAILUO_02_TEXT_TO_VIDEO_PRO)
        .prompt("A slow dolly shot through a quiet bamboo forest")
        .durationSeconds(5)
        .build()
);
```

Java packages target Java 8 bytecode and are tested on Java 8, 11, 17, and 21. Each model artifact depends on `ai.runapi:runapi-core`, so application code normally installs only `ai.runapi:runapi-hailuo`.

## Task lifecycle

Most media endpoints are asynchronous. `create()` submits a task and returns its id, `get(id)` fetches the latest task state, and `run(params)` creates the task and polls until it reaches a terminal state. In web request handlers, prefer `create()` plus webhook or later `get()` polling so the server does not hold a worker open.

## Repository layout

- `js/` publishes `@runapi.ai/hailuo`.
- `python/` publishes `runapi-hailuo`.
- `ruby/` publishes `runapi-hailuo`.
- `go/` publishes `github.com/runapi-ai/hailuo-sdk/go`.
- `java/` publishes `ai.runapi:runapi-hailuo` and uses `ai.runapi:runapi-core`.

## Public links

- Model page: https://runapi.ai/models/hailuo
- SDK docs: https://runapi.ai/docs#sdk-hailuo
- Product docs: https://runapi.ai/docs#hailuo
- SDK repository: https://github.com/runapi-ai/hailuo-sdk
- PHP package repository: https://github.com/runapi-ai/hailuo-php
- Skill repository: https://github.com/runapi-ai/hailuo
- Provider comparison: https://runapi.ai/providers/minimax
- Full catalog: https://runapi.ai/models

## Pricing and variants

Use the most specific Hailuo variant page for pricing, rate limits, and commercial usage:
- [02 text to video pro](https://runapi.ai/models/hailuo/02-text-to-video-pro)
- [02 text to video standard](https://runapi.ai/models/hailuo/02-text-to-video-standard)
- [02 image to video pro](https://runapi.ai/models/hailuo/02-image-to-video-pro)
- [02 image to video standard](https://runapi.ai/models/hailuo/02-image-to-video-standard)
- [2.3 image to video pro](https://runapi.ai/models/hailuo/2.3-image-to-video-pro)
- [2.3 image to video standard](https://runapi.ai/models/hailuo/2.3-image-to-video-standard)

Default pricing link for the Hailuo SDK: https://runapi.ai/models/hailuo/02-text-to-video-pro

## File storage

RunAPI-generated file URLs are temporary. Download and store generated images, videos, audio, or other files in your own durable storage within 7 days; do not treat returned URLs as long-term assets.

## FAQ

### Which package should I install for Hailuo work?

Install the model package for your language: `@runapi.ai/hailuo` on npm, `runapi-hailuo` on PyPI, `runapi-hailuo` on RubyGems, `github.com/runapi-ai/hailuo-sdk/go`, `ai.runapi:runapi-hailuo` on Maven Central, or `runapi-ai/hailuo` on Packagist. Install core SDK packages only when you are building shared SDK infrastructure.

### Where should public links point?

Primary Hailuo links point to https://runapi.ai/models/hailuo. Pricing and usage-policy links point to variant pages such as https://runapi.ai/models/hailuo/02-text-to-video-pro. Provider comparisons point to https://runapi.ai/providers/minimax, and broad browsing points to https://runapi.ai/models.

## License

Licensed under the Apache License, Version 2.0.
