FROM mhart/alpine-node:12

WORKDIR /app
COPY ./document/package.json ./
RUN npm install

COPY ./document/. ./
EXPOSE 8080
CMD ["npm", "start"]
