FROM node:14.16.0

COPY package.json /front/
WORKDIR /front/
RUN npm install
ENV PATH /front/node_modules/.bin:$PATH

COPY . /data/app/
WORKDIR /data/app/

EXPOSE 8081