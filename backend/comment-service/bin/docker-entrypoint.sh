#!/bin/bash -e

set -e

bundle exec rake db:prepare
# bundle exec rake db:drop db:create db:migrate db:seed

exec "${@}"
