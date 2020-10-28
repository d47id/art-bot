FROM golang:1.15 AS build

WORKDIR /go/src/app
COPY . /go/src/app

RUN go build -ldflags "-s -w" -o /go/bin/app

FROM gcr.io/distroless/base
COPY --from=build-env /go/bin/app /
ENTRYPOINT [ "/app" ]