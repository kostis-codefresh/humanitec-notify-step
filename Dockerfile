FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/workdir

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM alpine:3.12

RUN apk add --no-cache jq ca-certificates

COPY --from=build-env /tmp/workdir/humanitec-notify-step /app/humanitec-notify-step

WORKDIR /app

CMD ["./humanitec-notify-step"]
