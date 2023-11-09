FROM ruby:3.2.2-slim

WORKDIR /app

RUN apt-get update -qq && \
    apt-get install --no-install-recommends -y build-essential


COPY Gemfile Gemfile.lock ./
RUN bundle install

EXPOSE 3008

CMD ["ruby", "main.rb"]
