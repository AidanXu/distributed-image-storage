FROM node:20.11.1

WORKDIR /app

COPY . .

RUN rm -rf node_modules && npm install

EXPOSE 3000

CMD ["npm", "start"]
