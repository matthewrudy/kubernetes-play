#!/usr/bin/env ruby
require 'sinatra'

SECRETS_DIR = "/etc/secrets"

get '/' do
  if Dir.exist?("/etc/secrets")
    "Secrets:\n#{Dir["/etc/secrets/**/*"]}"
  else
    "No secrets"
  end
end
