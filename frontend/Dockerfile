FROM node:14

RUN mkdir /node && \
mkdir /node/src && \
mkdir /node/src/work

WORKDIR /node/src/work

ADD . /node/src/work
RUN yarn install

ENTRYPOINT ["yarn", "start"]
