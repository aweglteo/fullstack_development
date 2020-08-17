FROM ruby:latest

RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs
RUN apt-get install -y qt5-default qt5-qmake libqt5webkit5-dev gstreamer1.0-plugins-base gstreamer1.0-tools gstreamer1.0-x

RUN mkdir /work & mkdir /work/migration
WORKDIR /work/migration

ADD ./ /work/migration
RUN bundle install

CMD ["bundle", "exec", "ridgepole", "-c", "config.yml", "--apply"]
