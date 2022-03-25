FROM node:12

WORKDIR /js_parser
COPY . .
RUN apt update &&\
    apt-get install -y libnss3 libnss3-dev libgdk-pixbuf2.0-dev libgtk-3-dev libxss-dev libasound2 
RUN npm install
COPY . .
RUN ls