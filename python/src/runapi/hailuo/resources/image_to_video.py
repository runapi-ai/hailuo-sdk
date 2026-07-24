"""Hailuo image-to-video resource."""

from __future__ import annotations

from typing import Any, Dict, Optional

from runapi.core import Resource, ValidationError, RequestOptions

from ..contract_gen import CONTRACT
from ..types import (
    CompletedVideoTaskResponse,
    VideoTaskResponse,
)


class ImageToVideo(Resource):
    """Generate videos from a source image with Hailuo models."""

    ENDPOINT = "/api/v1/hailuo/image_to_video"

    RESPONSE_CLASS = VideoTaskResponse
    COMPLETED_RESPONSE_CLASS = CompletedVideoTaskResponse

    def run(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create an image-to-video task and poll until it completes.

        Args:
            **params: Image-to-video parameters (model, prompt, ...).

        Returns:
            The completed task with videos.
        """
        task = self.create(options=options, **params)
        return self._poll_until_complete(lambda: self.get(task.id, options=options))

    def create(self, options: Optional[RequestOptions] = None, **params: Any) -> Any:
        """Create an image-to-video task and return immediately with an ``id``.

        Args:
            **params: Image-to-video parameters (model, prompt, ...).

        Returns:
            The task creation result with an id.
        """
        compacted = self._compact_params(params)
        self._validate_params(compacted)
        return self._request("post", self.ENDPOINT, body=compacted, options=options)

    def get(self, id: str, options: Optional[RequestOptions] = None) -> Any:
        """Fetch the current status of an image-to-video task.

        Args:
            id: The task id.

        Returns:
            The current task status.
        """
        return self._request("get", f"{self.ENDPOINT}/{id}", options=options)

    def _validate_params(self, params: Dict[str, Any]) -> None:
        self._validate_contract(CONTRACT["image-to-video"], params)

        if not params.get("prompt"):
            raise ValidationError("prompt is required")

        model = params.get("model")
        duration_seconds = params.get("duration_seconds")
        output_resolution = params.get("output_resolution")

        if model == "hailuo-02-image-to-video-pro":
            if duration_seconds:
                raise ValidationError(f"duration_seconds is not supported for {model}")
            if output_resolution:
                raise ValidationError(f"output_resolution is not supported for {model}")
        elif model in ("hailuo-2.3-image-to-video-pro", "hailuo-2.3-image-to-video-standard"):
            if params.get("last_frame_image_url"):
                raise ValidationError(f"last_frame_image_url is not supported for {model}")
            if params.get("prompt_optimizer"):
                raise ValidationError(f"prompt_optimizer is not supported for {model}")
            if duration_seconds == 10 and str(output_resolution) == "1080p":
                raise ValidationError("1080p does not support 10-second duration")
