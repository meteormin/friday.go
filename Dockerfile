FROM alpine AS base

LABEL maintainer="meteormin"

ARG TIME_ZONE="Asia/Seoul"

RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/${TIME_ZONE} /etc/localtime && \
	echo "${TIME_ZONE}" > /etc/timezone \
	apk del tzdata


FROM node:22.13-alpine AS ui

WORKDIR /app

COPY ui .

RUN yarn install && yarn build

FROM golang:1.23-alpine AS build

RUN apk add --no-cache make gcc musl-dev

ARG MOD=friday

WORKDIR /app

COPY . .

COPY --from=ui /app/dist ./ui/dist

RUN go mod download && go build ./cmd/$MOD/main.go

FROM base AS  deploy

WORKDIR /app

COPY --from=build /app/main .
COPY --from=build /app/config.yml .

RUN mkdir -p /app/data

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "./main"]