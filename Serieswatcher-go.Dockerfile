FROM golang:1.11.3-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git curl
RUN curl -fsSL -o /usr/local/bin/dep \
    https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep

COPY . /go/src/gitea.fge.cloud/fabian_gehrlicher/series-watcher-v3
WORKDIR /go/src/gitea.fge.cloud/fabian_gehrlicher/series-watcher-v3

RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

RUN go build .

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/fgehrlicher/series-monitoringconfig.json /config.json
COPY --from=builder /app /app

ENV SUBDOMAIN gitea.fge.cloud
ENV PORT 443
ENV REVERSE_PROXY_URL 443

COPY ./register_subdomain.sh /register_subdomain
RUN chmod +x /register_subdomain

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/app"]