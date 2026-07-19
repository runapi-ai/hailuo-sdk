# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::Hailuo::Resources::TextToVideo do
  let(:http) { instance_double(RunApi::Core::HttpClient) }
  let(:resource) { described_class.new(http) }
  let(:endpoint) { "/api/v1/hailuo/text_to_video" }

  it "posts happy path" do
    params = {model: "hailuo-02-text-to-video-standard", prompt: "A drifting camera"}
    expect(http).to receive(:request).with(:post, endpoint, body: params).and_return("id" => "task-t2v-1")

    result = resource.create(**params)
    expect(result.id).to eq("task-t2v-1")
  end

  it "rejects invalid duration_seconds for pro model" do
    expect {
      resource.create(model: "hailuo-02-text-to-video-pro", prompt: "x", duration_seconds: 6)
    }.to raise_error(RunApi::Core::ValidationError, /duration_seconds is not supported/)
  end
end
