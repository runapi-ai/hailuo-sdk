# frozen_string_literal: true

module RunApi
  module Hailuo
    module Types
      TEXT_TO_VIDEO_MODELS = %w[hailuo-02-text-to-video-pro hailuo-02-text-to-video-standard].freeze
      IMAGE_TO_VIDEO_MODELS = %w[
        hailuo-02-image-to-video-pro
        hailuo-02-image-to-video-standard
        hailuo-2.3-image-to-video-pro
        hailuo-2.3-image-to-video-standard
      ].freeze
      DURATIONS = %w[6 10].freeze
      IMAGE_02_RESOLUTIONS = %w[512P 768P].freeze
      IMAGE_23_RESOLUTIONS = %w[768P 1080P].freeze

      class MediaUrl < RunApi::Core::BaseModel
        optional :url, String
      end

      class VideoTaskResponse < RunApi::Core::TaskResponse
        required :id, String
        optional :status, String, enum: -> { RunApi::Core::TaskResponse::Status::ALL }
        optional :videos, [ -> { MediaUrl } ]
        optional :error, String
      end

      class CompletedVideoTaskResponse < VideoTaskResponse
        required :videos, [ -> { MediaUrl } ]
      end
    end
  end
end
