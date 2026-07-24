# frozen_string_literal: true

module RunApi
  module Hailuo
    module Resources
      # Hailuo image-to-video resource.
      # Animate a still image into video guided by a text prompt and first-frame image.
      class ImageToVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/hailuo/image_to_video"

        RESPONSE_CLASS = Types::VideoTaskResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedVideoTaskResponse

        def initialize(http)
          @http = http
        end

        # Generate an image-to-video task and wait until complete.
        #
        # @param params [Hash] image-to-video parameters
        # @return [RunApi::Hailuo::Types::CompletedVideoTaskResponse] completed task with videos
        def run(options: nil, **params)
          task = create(options: options, **params)
          poll_until_complete { get(task.id, options: options) }
        end

        # Create an image-to-video task.
        #
        # @param params [Hash] image-to-video parameters
        # @return [RunApi::Hailuo::Types::VideoTaskResponse] task creation result with id
        def create(options: nil, **params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params, options: options)
        end

        # Get image-to-video task status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::Hailuo::Types::VideoTaskResponse] current task status
        def get(id, options: nil)
          request(:get, "#{ENDPOINT}/#{id}", options: options)
        end

        private

        def validate_params!(params)
          validate_contract!(CONTRACT["image-to-video"], params)

          raise Core::ValidationError, "prompt is required" unless param(params, :prompt)

          model = param(params, :model)
          duration_seconds = param(params, :duration_seconds)
          output_resolution = param(params, :output_resolution)

          case model
          when "hailuo-02-image-to-video-pro"
            raise Core::ValidationError, "duration_seconds is not supported for #{model}" if duration_seconds
            raise Core::ValidationError, "output_resolution is not supported for #{model}" if output_resolution
          when "hailuo-2.3-image-to-video-pro", "hailuo-2.3-image-to-video-standard"
            raise Core::ValidationError, "last_frame_image_url is not supported for #{model}" if param(params, :last_frame_image_url)
            raise Core::ValidationError, "prompt_optimizer is not supported for #{model}" if param(params, :prompt_optimizer)
            if duration_seconds == 10 && output_resolution.to_s == "1080p"
              raise Core::ValidationError, "1080p does not support 10-second duration"
            end
          end
        end
      end
    end
  end
end
