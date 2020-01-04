FROM golang:1.13
RUN mkdir src/dnsServer
ADD . src/dnsServer
WORKDIR src/dnsServer

#RUN chmod 4777 /go/src/

RUN go get "gopkg.in/yaml.v2"
RUN go get "github.com/jinzhu/gorm"
RUN go get "github.com/jinzhu/gorm/dialects/sqlite"

RUN go build -o main .
CMD ["/go/src/dnsServer/main"]