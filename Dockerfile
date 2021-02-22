FROM golang:1.15.1-alpine3.12 AS build-env

WORKDIR /tmp/workdir

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build 

FROM gcr.io/distroless/base
COPY --from=build-env /tmp/workdir/humanitec-notify-step /app/humanitec-notify-step

WORKDIR /app

CMD ["./humanitec-notify-step"]
