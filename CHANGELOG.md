# Changelog

## [js/v0.3.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.3.0), [go/v0.3.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.3.0) - 2026-09-29

### Changed
- Reject a 10-second duration at 1080p for hailuo-2.3-image-to-video-pro and hailuo-2.3-image-to-video-standard through generated contract rules, matching the API.
  Migration: Use 768p for 10-second Hailuo 2.3 videos, or a 6-second duration at 1080p.
- Record the server default duration and resolution of Hailuo models in generated contract metadata.

## [ruby/v0.3.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.3.0), [python/v0.3.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.3.0) - 2026-09-29

### Changed
- Reject a 10-second duration at 1080p for hailuo-2.3-image-to-video-pro and hailuo-2.3-image-to-video-standard through generated contract rules, matching the API.
  Migration: Use 768p for 10-second Hailuo 2.3 videos, or a 6-second duration at 1080p.
- Record the server default duration and resolution of Hailuo models in generated contract metadata.
- Enforce the Hailuo 2.3 1080p 10-second rule only through the generated contract rules; the rejection and its message are unchanged.


## [js/v0.2.11](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.11), [go/v0.2.11](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.11) - 2026-09-28

### Added
- Return usage.cost as a float USD amount on completed async Task query and webhook envelopes.

### Removed
- Remove the public Task billing object from Task envelopes.
  Migration: Read usage.cost on completed Task envelopes. Create, processing, and failed envelopes omit usage.


## [ruby/v0.2.12](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.12) - 2026-09-04

### Changed
- Update the `runapi-core` dependency range so this package remains installable with other current RunAPI Ruby SDKs.


## [ruby/v0.2.11](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.11) - 2026-08-18

### Changed
- Allow Ruby clients to install the core SDK release that adds persistent Files and multipart Uploads alongside this model SDK.


## [js/v0.2.10](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.10), [ruby/v0.2.10](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.10), [go/v0.2.10](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.10), [python/v0.2.3](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.2.3) - 2026-08-14

### Fixed
- Accept enable_safety_checker for Hailuo 02 Pro image-to-video requests.


## [js/v0.2.9](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.9), [ruby/v0.2.9](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.9), [go/v0.2.9](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.9), [python/v0.2.2](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.2.2) - 2026-08-12

### Fixed
- Reject enable_safety_checker for Hailuo 02 Pro image-to-video requests while preserving it for supported Hailuo models.


## [python/v0.2.1](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.2.1) - 2026-07-29

### Fixed
- Point package documentation metadata to the current RunAPI Developer Docs.


## [go/v0.2.8](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.8) - 2026-07-28

### Added
- Expose persisted billing facts on task responses.

## [js/v0.2.8](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.8) - 2026-07-28

### Added
- Type task billing facts on task responses.

## [ruby/v0.2.8](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.8) - 2026-07-28

### Added
- Expose live pricing through the shared core SDK.


## [python/v0.2.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.2.0) - 2026-07-24

### Added
- Expose shared Files, Account, and Pricing resources plus typed Task Billing Facts through the Provider Client.


## [js/v0.2.7](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.7), [ruby/v0.2.7](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.7), [go/v0.2.7](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.7) - 2026-07-02

### Fixed
- Request validation now derives allowed values (aspect ratios, output resolutions, formats) from the RunAPI request contract, so valid requests are no longer rejected client-side.
- Corrected field names and widened enum coverage for image generation endpoints.
- Documented reference image URL parameters where supported.

## [java/v0.1.1](https://github.com/runapi-ai/hailuo-sdk/releases/tag/java%2Fv0.1.1) - 2026-06-25

### Fixed
- Fixed Java retry handling for Retry-After response headers.
- Fixed Java contract validation for action-level conditional rules.
- Refreshed Java SDK metadata for v0.1.1.

## [java/v0.1.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/java%2Fv0.1.0) - 2026-06-24

### Added
- Publish `ai.runapi:runapi-hailuo` for Java SDK consumers.
- Include typed Java builders, synchronous client resources, sources, and Javadocs.

## [js/v0.2.6](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.6), [ruby/v0.2.6](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.6), [go/v0.2.6](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.6), [python/v0.1.0](https://github.com/runapi-ai/hailuo-sdk/releases/tag/python%2Fv0.1.0) - 2026-06-18

### Changed
- Per-method documentation for all resource methods

## [js/v0.2.5](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.5), [ruby/v0.2.5](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.5), [go/v0.2.5](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.5) - 2026-06-01

### Changed
- Align SDK with upstream Input Contract and public API vocabulary changes
- Update endpoint definitions and field constraints

## [js/v0.2.4](https://github.com/runapi-ai/hailuo-sdk/releases/tag/js%2Fv0.2.4), [ruby/v0.2.4](https://github.com/runapi-ai/hailuo-sdk/releases/tag/ruby%2Fv0.2.4), [go/v0.2.4](https://github.com/runapi-ai/hailuo-sdk/releases/tag/go%2Fv0.2.4) - 2026-05-22

### Changed
- Publish JavaScript, Ruby, and Go SDK artifacts for hailuo with per-language GitHub release tags.
- Refresh public README metadata.

## [v0.2.3](https://github.com/runapi-ai/hailuo-sdk/releases/tag/v0.2.3) - 2026-05-22

### Changed
- Publish hailuo-sdk v0.2.3 with refreshed README header, package metadata, and current SDK source.

## [v0.2.1](https://github.com/runapi-ai/hailuo-sdk/releases/tag/v0.2.1) - 2026-05-19

Initial release.
