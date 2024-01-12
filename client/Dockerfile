FROM node:20-alpine AS build

ENV PUBLIC_API_URL=http://localhost:1323
WORKDIR /app

COPY . .
RUN npm install
RUN npm run build

EXPOSE 3000

CMD ["node", "build"]
