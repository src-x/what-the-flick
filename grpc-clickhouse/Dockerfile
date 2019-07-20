FROM golang
RUN mkdir /go/src/grpc-clickhouse
ADD . /go/src/grpc-clickhouse
WORKDIR /go/src/grpc-clickhouse
RUN go get -u -v github.com/gin-contrib/cors
RUN go get -u -v github.com/gin-gonic/gin
RUN go get -u -v google.golang.org/grpc
RUN go get -u -v github.com/jmoiron/sqlx
RUN go get -u -v github.com/kshvakov/clickhouse
RUN go build ./recommend_client/client.go
RUN go build ./recommend_server/server.go
CMD ["sh","-c", "/go/src/grpc-clickhouse/server & /go/src/grpc-clickhouse/client"]