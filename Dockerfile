FROM golang:1.9.1

COPY . /go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3

RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/fatih/color
RUN go get github.com/PuerkitoBio/goquery
RUN go get github.com/common-nighthawk/go-figure
RUN go get github.com/gorilla/handlers

WORKDIR /go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3
RUN go build .
CMD ["/go/src/bitbucket.org/fabian_gehrlicher/series-watcher-v3/series-watcher-v3"]
