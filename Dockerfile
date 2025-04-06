FROM golang:1.24 AS build
WORKDIR /go/src/
COPY . /go/src/
RUN go build -ldflags "-s -w" -o /go/bin/art-bot

FROM gcr.io/distroless/base
COPY --from=build /go/bin/art-bot /
ENTRYPOINT [ "/art-bot" ]