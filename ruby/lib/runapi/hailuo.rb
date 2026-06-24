# frozen_string_literal: true

require "runapi/core"
require_relative "hailuo/types"
require_relative "hailuo/contract_gen"
require_relative "hailuo/resources/text_to_video"
require_relative "hailuo/resources/image_to_video"
require_relative "hailuo/client"

module RunApi
  module Hailuo
    AuthenticationError = RunApi::Core::AuthenticationError
    RateLimitError = RunApi::Core::RateLimitError
    InsufficientCreditsError = RunApi::Core::InsufficientCreditsError
    NotFoundError = RunApi::Core::NotFoundError
    ValidationError = RunApi::Core::ValidationError
    TaskFailedError = RunApi::Core::TaskFailedError
    TaskTimeoutError = RunApi::Core::TaskTimeoutError
  end
end
