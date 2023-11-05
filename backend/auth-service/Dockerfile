FROM golang:1.21.3-alpine AS dependencies
WORKDIR /dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM dependencies AS build
WORKDIR /build
COPY . .
RUN go build -v -o ./service ./cmd/app

FROM alpine:3.18.4
WORKDIR /service
COPY --from=build /build/key ./key
COPY --from=build /build/service ./
CMD ["./service"]