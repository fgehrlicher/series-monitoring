FROM golang:1.11.3-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git curl
RUN curl -fsSL -o /usr/local/bin/dep \
    https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && \
    chmod +x /usr/local/bin/dep

COPY . /go/src/github.com/fgehrlicher/series-monitoring
WORKDIR /go/src/github.com/fgehrlicher/series-monitoring

RUN dep ensure -vendor-only
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

RUN go build .

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/fgehrlicher/series-monitoring/config.json /config.json
COPY --from=builder /app /app

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/app"]