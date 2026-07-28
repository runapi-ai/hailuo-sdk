# frozen_string_literal: true

require "spec_helper"

RSpec.describe RunApi::Hailuo::Client do
  before do
    allow(ConnectionPool).to receive(:new).and_return(instance_double(ConnectionPool))
  end

  after { RunApi.api_key = nil }

  it "exposes both resource accessors" do
    client = described_class.new(api_key: "test-key")
    expect(client.text_to_video).to be_a(RunApi::Hailuo::Resources::TextToVideo)
    expect(client.image_to_video).to be_a(RunApi::Hailuo::Resources::ImageToVideo)
  end
end
