FROM golang:1.14

ENV MYSQL_DB_HOST 192.168.1.32:3306
ENV MYSQL_DB kutuphane
ENV MYSQL_DB_USER burak
ENV MYSQL_DB_PASS 1234
WORKDIR /go/src/app
COPY . .
RUN go get -u gorm.io/gorm

RUN go get -d -v ./...
RUN go install -v ./...

CMD ["app"]