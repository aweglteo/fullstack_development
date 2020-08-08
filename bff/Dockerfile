FROM node:14

RUN mkdir /bff && \
mkdir /bff/src && \
mkdir /bff/src/work

WORKDIR /bff/src/work

ADD . /bff/src/work
RUN npm install

CMD ["npm", "run", "webpack"]
ENTRYPOINT ["npm", "run", "dev"]
