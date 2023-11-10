#!/bin/bash -e

set -e

bundle exec rake db:prepare

exec "${@}"
