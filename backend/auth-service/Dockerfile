FROM golang:1.21.3-alpine AS dependencies
WORKDIR /dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

FROM dependencies AS build
WORKDIR /build
COPY . .
RUN go build -v -o ./service ./cmd/app

FROM alpine:3.18.4
WORKDIR /app
RUN mkdir -p ./keys
RUN apk add --no-cache --update openssh-keygen && \
    ssh-keygen -t ecdsa -f ./keys/ecdsa -b 521 -m pem && \
    ssh-keygen -f ./keys/ecdsa -e -m pem > ./keys/ecdsa.pub
ENV KEY_PATH=./keys/ecdsa
COPY --from=build /build/service ./
CMD ["./service"]