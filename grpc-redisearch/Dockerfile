FROM golang
RUN mkdir /go/src/grpc-redisearch
ADD . /go/src/grpc-redisearch
WORKDIR /go/src/grpc-redisearch
RUN go get -v -u github.com/gin-contrib/cors
RUN go get -v -u github.com/gin-gonic/gin
RUN go get -v -u google.golang.org/grpc
RUN go get -v -u github.com/RediSearch/redisearch-go/redisearch
RUN go build ./search_client/client.go
RUN go build ./search_server/server.go
CMD ["sh", "-c", "/go/src/grpc-redisearch/server & /go/src/grpc-redisearch/client"]