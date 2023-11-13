FROM node:lts-alpine AS build
WORKDIR /build
COPY package*.json .
RUN npm install
COPY . .
RUN npm run build

FROM node:lts-alpine AS production
WORKDIR /app
COPY --from=build /build/dist ./dist
RUN npm install -g vite
CMD ["vite", "preview", "--host", "--port", "5173"]