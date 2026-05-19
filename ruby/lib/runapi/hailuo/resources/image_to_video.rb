# frozen_string_literal: true

module RunApi
  module Hailuo
    module Resources
      class ImageToVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/hailuo/image_to_video"

        RESPONSE_CLASS = Types::VideoTaskResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedVideoTaskResponse

        def initialize(http)
          @http = http
        end

        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          model = param(params, :model)
          raise Core::ValidationError, "model is required" unless Types::IMAGE_TO_VIDEO_MODELS.include?(model)
          raise Core::ValidationError, "prompt is required" unless param(params, :prompt)
          raise Core::ValidationError, "image_url is required" unless param(params, :image_url)

          duration = param(params, :duration)
          resolution = param(params, :resolution)

          if duration && !Types::DURATIONS.include?(duration.to_s)
            raise Core::ValidationError, "duration must be one of: #{Types::DURATIONS.join(", ")}"
          end

          case model
          when "hailuo-02-image-to-video-pro"
            raise Core::ValidationError, "duration is not supported for #{model}" if duration
            raise Core::ValidationError, "resolution is not supported for #{model}" if resolution
          when "hailuo-02-image-to-video-standard"
            validate_resolution!(resolution, Types::IMAGE_02_RESOLUTIONS) if resolution
          else
            validate_resolution!(resolution, Types::IMAGE_23_RESOLUTIONS) if resolution
            raise Core::ValidationError, "end_image_url is not supported for #{model}" if param(params, :end_image_url)
            raise Core::ValidationError, "prompt_optimizer is not supported for #{model}" if param(params, :prompt_optimizer)
            if duration.to_s == "10" && resolution.to_s == "1080P"
              raise Core::ValidationError, "1080P does not support 10-second duration"
            end
          end
        end

        def validate_resolution!(resolution, allowed)
          return if allowed.include?(resolution.to_s)

          raise Core::ValidationError, "resolution must be one of: #{allowed.join(", ")}"
        end
      end
    end
  end
end
