# frozen_string_literal: true

module RunApi
  module Hailuo
    module Resources
      # Hailuo text-to-video resource.
      # Generate video from a text prompt.
      class TextToVideo
        include RunApi::Core::ResourceHelpers

        ENDPOINT = "/api/v1/hailuo/text_to_video"

        RESPONSE_CLASS = Types::VideoTaskResponse
        COMPLETED_RESPONSE_CLASS = Types::CompletedVideoTaskResponse

        def initialize(http)
          @http = http
        end

        # Generate a text-to-video task and wait until complete.
        #
        # @param params [Hash] text-to-video parameters
        # @return [RunApi::Hailuo::Types::CompletedVideoTaskResponse] completed task with videos
        def run(**params)
          task = create(**params)
          poll_until_complete { get(task.id) }
        end

        # Create a text-to-video task.
        #
        # @param params [Hash] text-to-video parameters
        # @return [RunApi::Hailuo::Types::VideoTaskResponse] task creation result with id
        def create(**params)
          params = compact_params(params)
          validate_params!(params)
          request(:post, ENDPOINT, body: params)
        end

        # Get text-to-video task status by task ID.
        #
        # @param id [String] task ID
        # @return [RunApi::Hailuo::Types::VideoTaskResponse] current task status
        def get(id)
          request(:get, "#{ENDPOINT}/#{id}")
        end

        private

        def validate_params!(params)
          model = param(params, :model)
          raise Core::ValidationError, "model is required" unless Types::TEXT_TO_VIDEO_MODELS.include?(model)
          raise Core::ValidationError, "prompt is required" unless param(params, :prompt)

          duration_seconds = param(params, :duration_seconds)
          return unless duration_seconds

          unless Types::DURATIONS.include?(duration_seconds)
            raise Core::ValidationError, "duration_seconds must be one of: #{Types::DURATIONS.join(", ")}"
          end
          if model == "hailuo-02-text-to-video-pro"
            raise Core::ValidationError, "duration_seconds is not supported for #{model}"
          end
        end
      end
    end
  end
end
