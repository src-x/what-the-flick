package main

import (
	"fmt"
	"grpc-clickhouse/recommendpb"
	"log"
	"net"

	"github.com/jmoiron/sqlx"
	_ "github.com/kshvakov/clickhouse"
	"google.golang.org/grpc"
)

func dbQuery(movieName string) []struct {
	ID        int64   "db:\"id\""
	Title     string  "db:\"title\""
	CosineSim float64 "db:\"cosinesim\""
} {
	connect, err := sqlx.Open("clickhouse", "tcp://clickhouse:9000?debug=true")
	if err != nil {
		log.Fatal(err)
	}
	var items []struct {
		ID        int64   `db:"id"`
		Title     string  `db:"title"`
		CosineSim float64 `db:"cosinesim"`
	}
	movQuery := fmt.Sprintf("SELECT id, title, sum(ratsum)/(sqrt(sum(rata)) * sqrt(sum(ratb))) as cosinesim FROM (SELECT a.tmdbId AS id, a.title AS title, a.rating * b.rating AS ratsum, pow(a.rating, 2) as rata, pow(b.rating, 2) AS ratb FROM (SELECT * FROM test WHERE title IN (SELECT title FROM test WHERE userId IN (SELECT userId FROM test WHERE title = '%[1]s') GROUP BY title HAVING COUNT(*)>40) AND userId IN (SELECT userId FROM test WHERE title = '%[1]s')) AS a ALL INNER JOIN (SELECT * FROM test WHERE title = '%[1]s') as b ON a.userId = b.userId) WHERE title != '%[1]s' GROUP BY id, title ORDER BY cosinesim DESC LIMIT 10", movieName)
	if err := connect.Select(&items, movQuery); err != nil {
		log.Fatal(err)
	}
	return items
}

type server struct{}

func (*server) Recommend(in *recommendpb.RecommendRequest, stream recommendpb.RecommendService_RecommendServer) error {
	moviename := in.GetRecommender().GetMovie()
	result := dbQuery(moviename)
	for _, item := range result {
		res := &recommendpb.RecommendResponse{
			MovieTitle: fmt.Sprintf("%s", item.Title),
			MovieId:    item.ID,
		}
		stream.Send(res)
	}
	return nil
}

func main() {

	lis, err := net.Listen("tcp", "0.0.0.0:50050")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	recommendpb.RegisterRecommendServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
