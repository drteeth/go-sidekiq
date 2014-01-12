### Go job processor for Sidekiq.

## What does it do?
Pulls sidekiq jobs out of redis and runs them.

## How can it run the body of my ruby sidekiq jobs?!
It can't. You have to port your jobs to Go.

## What the eff use is it then?
Probably not much tbh. It lets you keep the same calling semantics in your ruby code.  Once finished it would also respect the queuing/priority semantics of sidekiq that you know and love.

## Usage:
Given a redis server at 127.0.0.1:6379

    # build it
    go build

    # stuff some tasks into sidekiq
    ruby app.rb
    ruby app.rb
    ruby app.rb
    ruby app.rb

    # start up the go job processor
    ./go-sidekiq

    # watch as go does work, queue some more jobs, stand back in amazement.
