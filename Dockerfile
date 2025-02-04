FROM alpine AS base

LABEL maintainer="meteormin"

ARG TIME_ZONE

RUN apk --no-cache add tzdata && \
	cp /usr/share/zoneinfo/Asia/Seoul /etc/localtime && \
	echo "${TIME_ZONE}" > /etc/timezone \
	apk del tzdata

FROM golang:1.22-alpine AS build

RUN apk add --no-cache make gcc musl-dev

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 go mod download
RUN make build

FROM base AS  deploy

WORKDIR /app

COPY --from=build /app/build/ .
COPY --from=build /app/build/cli/simul-cli .

EXPOSE 8080

ENTRYPOINT ["sh", "-c", "./ -port=8080"]