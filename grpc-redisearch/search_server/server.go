package main

import (
	"grpc-redisearch/redisearchpb"
	"log"
	"net"
	"strconv"

	"github.com/RediSearch/redisearch-go/redisearch"
	"google.golang.org/grpc"
)

func searchQuery(movieName string) []redisearch.Document {
	c := redisearch.NewClient("redisearch:6379", "movie")

	docs, _, _ := c.Search(redisearch.NewQuery(movieName).
		Limit(0, 10).
		SetReturnFields("title"))

	return docs
}

func autoCompleteQuery(movieName string) []redisearch.Suggestion {
	a := redisearch.NewAutocompleter("redisearch:6379", "ac")

	suggestions, _ := a.Suggest(movieName, 10, true)

	return suggestions

}

type server struct{}

func (*server) Suggest(in *redisearchpb.Request, stream redisearchpb.Service_SuggestServer) error {
	moviename := in.GetSuggester().GetMovie()
	result := autoCompleteQuery(moviename)
	for i := range result {
		res := &redisearchpb.SuggestResponse{
			Title: result[i].Term,
		}
		stream.Send(res)
	}
	return nil
}

func (*server) Search(in *redisearchpb.Request, stream redisearchpb.Service_SearchServer) error {
	moviename := in.GetSearcher().GetMovie()
	result := searchQuery(moviename)
	for i := range result {
		id, _ := strconv.ParseInt(result[i].Id, 10, 64)
		res := &redisearchpb.SearchResponse{
			Title:   result[i].Properties["title"].(string),
			MovieId: id,
		}
		stream.Send(res)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	redisearchpb.RegisterServiceServer(s, &server{})
	// recommendpb.RegisterRecommendServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
