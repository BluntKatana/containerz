FROM node:20-alpine AS build

WORKDIR /app

COPY . .
RUN npm install
RUN npm run build --

CMD ["node", "build"]
