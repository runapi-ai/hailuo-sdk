# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::Hailuo::Resources::ImageToVideo do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/hailuo/image_to_video" }

  it "posts happy path" do
    params = {
      model: "hailuo-2.3-image-to-video-standard",
      prompt: "Animate the portrait",
      first_frame_image_url: "https://cdn.runapi.ai/public/samples/input.png"
    }
    expect(http).to receive(:request).with(:post, endpoint, body: params).and_return("id" => "task-i2v-1")

    result = resource.create(**params)
    expect(result.id).to eq("task-i2v-1")
  end

  it "rejects 1080p with 10-second duration on 2.3 models" do
    expect {
      resource.create(
        model: "hailuo-2.3-image-to-video-standard",
        prompt: "Animate the portrait",
        first_frame_image_url: "https://cdn.runapi.ai/public/samples/input.png",
        duration_seconds: 10,
        output_resolution: "1080p"
      )
    }.to raise_error(RunApi::Core::ValidationError, /1080p does not support 10-second duration/)
  end
end
