#!/usr/bin/env ruby
require 'sinatra'

SECRETS_DIR = "/etc/secrets"

get '/' do
  if Dir.exist?(SECRETS_DIR)
    rtn = "Secrets:\n"
    Dir["#{SECRETS_DIR}/**/*"].each do |path|
      raw = File.read(path)
      decoded = raw.unpack("m")

      rtn << "  #{path}:"
      rtn << "    raw: #{raw}"
      rtn << "    decoded: #{decoded}"
    end
    rtn
  else
    "No secrets"
  end
end
