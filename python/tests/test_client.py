import pytest

from runapi.core import config
from runapi.core.errors import AuthenticationError
from runapi.hailuo import HailuoClient
from runapi.hailuo.resources.image_to_video import ImageToVideo
from runapi.hailuo.resources.text_to_video import TextToVideo
from runapi.hailuo.types import CompletedVideoTaskResponse, VideoTaskResponse


class FakeHttp:
    def __init__(self, *responses):
        self._responses = list(responses)
        self.calls = []

    def request(self, method, path, body=None, options=None):
        self.calls.append((method, path, body))
        if self._responses:
            return self._responses.pop(0)
        return {"id": "task_1", "status": "pending"}


@pytest.fixture(autouse=True)
def reset_config(monkeypatch):
    monkeypatch.delenv("RUNAPI_API_KEY", raising=False)
    monkeypatch.setattr(config, "api_key", None)
    yield


# --- auth -----------------------------------------------------------------


def test_accepts_api_key_parameter():
    assert isinstance(HailuoClient(api_key="k", http_client=FakeHttp()), HailuoClient)


def test_falls_back_to_global(monkeypatch):
    monkeypatch.setattr(config, "api_key", "global-key")
    assert isinstance(HailuoClient(http_client=FakeHttp()), HailuoClient)


def test_falls_back_to_env(monkeypatch):
    monkeypatch.setenv("RUNAPI_API_KEY", "env-key")
    assert isinstance(HailuoClient(http_client=FakeHttp()), HailuoClient)


def test_raises_without_api_key():
    with pytest.raises(AuthenticationError, match="API key is required"):
        HailuoClient()


# --- injection / accessors ------------------------------------------------


def test_uses_injected_http_client():
    fake = FakeHttp()
    client = HailuoClient(api_key="k", http_client=fake)
    assert client.text_to_video._http is fake
    assert client.image_to_video._http is fake


def test_exposes_resource_accessors():
    client = HailuoClient(api_key="k", http_client=FakeHttp())
    assert isinstance(client.text_to_video, TextToVideo)
    assert isinstance(client.image_to_video, ImageToVideo)


# --- request shapes -------------------------------------------------------


def test_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = HailuoClient(api_key="k", http_client=fake)
    result = client.text_to_video.create(
        model="hailuo-02-text-to-video-standard",
        prompt="hello world",
        duration_seconds=6,
        output_resolution=None,
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/hailuo/text_to_video",
            {"model": "hailuo-02-text-to-video-standard", "prompt": "hello world", "duration_seconds": 6},
        )]
    assert isinstance(result, VideoTaskResponse)


def test_image_to_video_create_posts_compacted_body():
    fake = FakeHttp({"id": "t1", "status": "pending"})
    client = HailuoClient(api_key="k", http_client=fake)
    client.image_to_video.create(
        model="hailuo-2.3-image-to-video-standard",
        prompt="pan the camera",
        first_frame_image_url="https://x/a.png",
        output_resolution="768p",
    )
    assert fake.calls == [
        (
            "post",
            "/api/v1/hailuo/image_to_video",
            {
                "model": "hailuo-2.3-image-to-video-standard",
                "prompt": "pan the camera",
                "first_frame_image_url": "https://x/a.png",
                "output_resolution": "768p"},
        )]


def test_get_fetches_by_id():
    fake = FakeHttp({"id": "t1", "status": "processing"})
    client = HailuoClient(api_key="k", http_client=fake)
    client.text_to_video.get("t1")
    assert fake.calls == [("get", "/api/v1/hailuo/text_to_video/t1", None)]


def test_run_narrows_completed_type():
    fake = FakeHttp(
        {"id": "t1", "status": "pending"},
        {"id": "t1", "status": "completed", "usage": {"cost": 0.05}, "videos": [{"url": "https://x/y.mp4"}]},
    )
    client = HailuoClient(api_key="k", http_client=fake)
    result = client.text_to_video.run(
        model="hailuo-02-text-to-video-standard", prompt="a serene lake"
    )
    assert isinstance(result, CompletedVideoTaskResponse)
    assert result.videos[0].url == "https://x/y.mp4"
