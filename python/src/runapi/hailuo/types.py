"""Hailuo model lists, enums, and response models."""

from __future__ import annotations

from runapi.core import BaseModel, TaskResponse, optional, required

TEXT_TO_VIDEO_MODELS = [
    "hailuo-02-text-to-video-pro",
    "hailuo-02-text-to-video-standard",
]
IMAGE_TO_VIDEO_MODELS = [
    "hailuo-02-image-to-video-pro",
    "hailuo-02-image-to-video-standard",
    "hailuo-2.3-image-to-video-pro",
    "hailuo-2.3-image-to-video-standard",
]
DURATIONS = [6, 10]
IMAGE_02_RESOLUTIONS = ["512p", "768p"]
IMAGE_23_RESOLUTIONS = ["768p", "1080p"]


class MediaUrl(BaseModel):
    url = optional(str)


class VideoTaskResponse(TaskResponse):
    """Response for a video generation task."""

    id = required(str)
    status = optional(str, enum=lambda: TaskResponse.Status.ALL)
    videos = optional([lambda: MediaUrl])
    error = optional(str)


class CompletedVideoTaskResponse(VideoTaskResponse):
    """Narrowed response from ``run()`` once polling observes completion."""

    videos = required([lambda: MediaUrl])
