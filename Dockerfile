FROM alpine AS base

LABEL maintainer="meteormin"

ARG TIME_ZONE="Asia/Seoul"

RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime && \
	echo "${TIME_ZONE}" > /etc/timezone \
	apk del tzdata

FROM golang:1.23-alpine AS build

RUN apk add --no-cache make gcc musl-dev

ARG MOD

WORKDIR /app

COPY . .

RUN go mod download && go build ./cmd/$MOD/main.go

FROM base AS  deploy

WORKDIR /app

COPY --from=build /app/build/ .

COPY config.yml .

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "./main -port=8080"]