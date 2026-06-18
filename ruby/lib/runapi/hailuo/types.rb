# frozen_string_literal: true

module RunApi
  module Hailuo
    # Type definitions and constants for Hailuo video generation.
    module Types
      # Text-to-video model variants. Pro produces higher-fidelity output;
      # Standard is faster with slightly lower quality. Prompts limited to 1500 characters.
      TEXT_TO_VIDEO_MODELS = %w[hailuo-02-text-to-video-pro hailuo-02-text-to-video-standard].freeze

      # Image-to-video model variants spanning two generations.
      # 02 models: last-frame image, prompt optimizer, 512p/768p, 1500-char prompts.
      # 2.3 models: 768p/1080p output, 5000-char prompts, no last-frame or prompt optimizer.
      IMAGE_TO_VIDEO_MODELS = %w[
        hailuo-02-image-to-video-pro
        hailuo-02-image-to-video-standard
        hailuo-2.3-image-to-video-pro
        hailuo-2.3-image-to-video-standard
      ].freeze

      # Video duration options in seconds.
      DURATIONS = [6, 10].freeze

      # Output resolutions for 02 generation image-to-video models.
      IMAGE_02_RESOLUTIONS = %w[512p 768p].freeze

      # Output resolutions for 2.3 generation image-to-video models. 1080p is not available with 10-second duration.
      IMAGE_23_RESOLUTIONS = %w[768p 1080p].freeze

      # A generated video file with a download URL.
      class MediaUrl < RunApi::Core::BaseModel
        optional :url, String
      end

      # Video generation task status and results. Poll until status reaches a terminal state.
      class VideoTaskResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :videos, [-> { MediaUrl }]
        optional :error, String
      end

      # Narrowed response returned after successful completion, guaranteeing videos is present.
      class CompletedVideoTaskResponse < VideoTaskResponse
        required :videos, [-> { MediaUrl }]
      end
    end
  end
end
