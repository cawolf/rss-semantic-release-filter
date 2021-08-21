FROM golang:1.17 as build-env

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY . /app

RUN go build

###

FROM gcr.io/distroless/base

EXPOSE 1323

ENV RSS_SEMANTIC_RELEASE_FILTER_LOG_STRUCTURED=true

COPY --from=build-env /app/rss-semantic-release-filter /

ENTRYPOINT ["/rss-semantic-release-filter"]

CMD ["feed", "-directory=/data"]
