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
          raise Core::ValidationError, "first_frame_image_url is required" unless param(params, :first_frame_image_url)

          duration_seconds = param(params, :duration_seconds)
          output_resolution = param(params, :output_resolution)

          if duration_seconds && !Types::DURATIONS.include?(duration_seconds)
            raise Core::ValidationError, "duration_seconds must be one of: #{Types::DURATIONS.join(", ")}"
          end

          case model
          when "hailuo-02-image-to-video-pro"
            raise Core::ValidationError, "duration_seconds is not supported for #{model}" if duration_seconds
            raise Core::ValidationError, "output_resolution is not supported for #{model}" if output_resolution
          when "hailuo-02-image-to-video-standard"
            validate_output_resolution!(output_resolution, Types::IMAGE_02_RESOLUTIONS) if output_resolution
          else
            validate_output_resolution!(output_resolution, Types::IMAGE_23_RESOLUTIONS) if output_resolution
            raise Core::ValidationError, "last_frame_image_url is not supported for #{model}" if param(params, :last_frame_image_url)
            raise Core::ValidationError, "prompt_optimizer is not supported for #{model}" if param(params, :prompt_optimizer)
            if duration_seconds == 10 && output_resolution.to_s == "1080p"
              raise Core::ValidationError, "1080p does not support 10-second duration"
            end
          end
        end

        def validate_output_resolution!(output_resolution, allowed)
          return if allowed.include?(output_resolution.to_s)

          raise Core::ValidationError, "output_resolution must be one of: #{allowed.join(", ")}"
        end
      end
    end
  end
end
