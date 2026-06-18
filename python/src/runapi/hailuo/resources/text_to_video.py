"""Hailuo text-to-video resource."""

from __future__ import annotations

from typing import Any, Dict

from runapi.core import Resource, ValidationError

from ..types import (
    DURATIONS,
    TEXT_TO_VIDEO_MODELS,
    CompletedVideoTaskResponse,
    VideoTaskResponse,
)


class TextToVideo(Resource):
    """Generate videos from text prompts with Hailuo models."""

    ENDPOINT = "/api/v1/hailuo/text_to_video"

    RESPONSE_CLASS = VideoTaskResponse
    COMPLETED_RESPONSE_CLASS = CompletedVideoTaskResponse

    def run(self, **params: Any) -> Any:
        """Create a text-to-video task and poll until it completes.

        Args:
            **params: Text-to-video parameters (model, prompt, ...).

        Returns:
            The completed task with videos.
        """
        task = self.create(**params)
        return self._poll_until_complete(lambda: self.get(task.id))

    def create(self, **params: Any) -> Any:
        """Create a text-to-video task and return immediately with an ``id``.

        Args:
            **params: Text-to-video parameters (model, prompt, ...).

        Returns:
            The task creation result with an id.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted)

    def get(self, id: str) -> Any:
        """Fetch the current status of a text-to-video task.

        Args:
            id: The task id.

        Returns:
            The current task status.
        """
        return self._request("get", f"{self.ENDPOINT}/{id}")

    def _validate_params(self, params: Dict[str, Any]) -> None:
        model = params.get("model")
        if model not in TEXT_TO_VIDEO_MODELS:
            raise ValidationError("model is required")
        if not params.get("prompt"):
            raise ValidationError("prompt is required")

        duration_seconds = params.get("duration_seconds")
        if not duration_seconds:
            return

        if duration_seconds not in DURATIONS:
            raise ValidationError(
                f"duration_seconds must be one of: {', '.join(str(d) for d in DURATIONS)}"
            )
        if model == "hailuo-02-text-to-video-pro":
            raise ValidationError(f"duration_seconds is not supported for {model}")
