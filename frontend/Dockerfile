FROM node:20.11.1

WORKDIR /app

COPY package.json package-lock.json ./

RUN rm -rf node_modules && npm install

COPY . .

EXPOSE 3000

CMD ["npm", "start"]
