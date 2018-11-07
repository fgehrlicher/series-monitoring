FROM golang:1.9.1-alpine AS builder

RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

RUN apk add --no-cache ca-certificates git

COPY . /go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/fatih/color
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/common-nighthawk/go-figure
RUN go get github.com/gorilla/handlers

WORKDIR /go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3

RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /app .

RUN go build .

FROM scratch AS final

COPY --from=builder /user/group /user/passwd /etc/
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3/config.json /config.json
COPY --from=builder /app /app

USER nobody:nobody

EXPOSE 8080

ENTRYPOINT ["/app"]