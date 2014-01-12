require 'sidekiq'

# If your client is single-threaded, we just need a single connection in our Redis connection pool
Sidekiq.configure_client do |config|
  config.redis = { :namespace => 'x', :size => 1, :url => 'redis://127.0.0.1:6379' }
end

class PlainOldRuby
  include Sidekiq::Worker

  def perform(how_hard="super hard", how_long=1)
    # Go will take over
  end
end

PlainOldRuby.perform_async "like a dog", 5