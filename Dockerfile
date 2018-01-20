FROM golang:1.9.1
RUN mkdir /app
RUN go get github.com/gorilla/mux
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/fatih/color
ADD . /app/
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
