FROM node:14

RUN mkdir /bff && \
mkdir /bff/src && \
mkdir /bff/src/work

WORKDIR /bff/src/work

COPY startup.sh .
COPY package.json . 

RUN npm install

RUN chmod 744 ./startup.sh

ENTRYPOINT ["./startup.sh"]
