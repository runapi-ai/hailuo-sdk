"""Hailuo image-to-video resource."""

from __future__ import annotations

from typing import Any, Dict, Sequence

from runapi.core import Resource, ValidationError

from ..types import (
    DURATIONS,
    IMAGE_02_RESOLUTIONS,
    IMAGE_23_RESOLUTIONS,
    IMAGE_TO_VIDEO_MODELS,
    CompletedVideoTaskResponse,
    VideoTaskResponse,
)


class ImageToVideo(Resource):
    """Generate videos from a source image with Hailuo models."""

    ENDPOINT = "/api/v1/hailuo/image_to_video"

    RESPONSE_CLASS = VideoTaskResponse
    COMPLETED_RESPONSE_CLASS = CompletedVideoTaskResponse

    def run(self, **params: Any) -> Any:
        """Create an image-to-video task and poll until it completes.

        Args:
            **params: Image-to-video parameters (model, prompt, ...).

        Returns:
            The completed task with videos.
        """
        task = self.create(**params)
        return self._poll_until_complete(lambda: self.get(task.id))

    def create(self, **params: Any) -> Any:
        """Create an image-to-video task and return immediately with an ``id``.

        Args:
            **params: Image-to-video parameters (model, prompt, ...).

        Returns:
            The task creation result with an id.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted)

    def get(self, id: str) -> Any:
        """Fetch the current status of an image-to-video task.

        Args:
            id: The task id.

        Returns:
            The current task status.
        """
        return self._request("get", f"{self.ENDPOINT}/{id}")

    def _validate_params(self, params: Dict[str, Any]) -> None:
        model = params.get("model")
        if model not in IMAGE_TO_VIDEO_MODELS:
            raise ValidationError("model is required")
        if not params.get("prompt"):
            raise ValidationError("prompt is required")
        if not params.get("first_frame_image_url"):
            raise ValidationError("first_frame_image_url is required")

        duration_seconds = params.get("duration_seconds")
        output_resolution = params.get("output_resolution")

        if duration_seconds and duration_seconds not in DURATIONS:
            raise ValidationError(
                f"duration_seconds must be one of: {', '.join(str(d) for d in DURATIONS)}"
            )

        if model == "hailuo-02-image-to-video-pro":
            if duration_seconds:
                raise ValidationError(f"duration_seconds is not supported for {model}")
            if output_resolution:
                raise ValidationError(f"output_resolution is not supported for {model}")
        elif model == "hailuo-02-image-to-video-standard":
            if output_resolution:
                self._validate_output_resolution(output_resolution, IMAGE_02_RESOLUTIONS)
        else:
            if output_resolution:
                self._validate_output_resolution(output_resolution, IMAGE_23_RESOLUTIONS)
            if params.get("last_frame_image_url"):
                raise ValidationError(f"last_frame_image_url is not supported for {model}")
            if params.get("prompt_optimizer"):
                raise ValidationError(f"prompt_optimizer is not supported for {model}")
            if duration_seconds == 10 and str(output_resolution) == "1080p":
                raise ValidationError("1080p does not support 10-second duration")

    @staticmethod
    def _validate_output_resolution(output_resolution: Any, allowed: Sequence[str]) -> None:
        if str(output_resolution) in allowed:
            return
        raise ValidationError(f"output_resolution must be one of: {', '.join(allowed)}")
