package main

import (
	"context"
	"fmt"
	"grpc-redisearch/redisearchpb"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type SearchMoviesData struct {
	Movies []SearchMovie `json:"Movie"`
}

type SuggestMoviesData struct {
	Movies []SuggestMovie `json:"Movie"`
}

type SearchMovie struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

type SuggestMovie struct {
	Title string `json:"title"`
}

func main() {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := redisearchpb.NewServiceClient(cc)
	redisStream(c)
}
func redisStream(c redisearchpb.ServiceClient) {
	fmt.Println("starting rpc stream")
	g := gin.Default()
	g.Use(cors.Default())

	g.GET("/search/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		in := &redisearchpb.Request{
			Searcher: &redisearchpb.Searcher{
				Movie: id,
			},
		}
		var mainStruct SearchMoviesData
		if res, err := c.Search(context.Background(), in); err == nil {
			for {
				msg, err := res.Recv()
				if err == io.EOF {
					// we've reached the end of the stream
					break
				}
				if err != nil {
					log.Fatalf("search error: %v", err)
				}
				log.Printf("Search Response: %v", msg.GetTitle())
				currentMovie := SearchMovie{msg.GetMovieId(), msg.GetTitle()}
				mainStruct.Movies = append(mainStruct.Movies, currentMovie)
			}
			ctx.JSON(http.StatusOK, mainStruct)
		}
	})
	g.GET("/suggest/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		in := &redisearchpb.Request{
			Suggester: &redisearchpb.Suggester{
				Movie: id,
			},
		}
		var mainStruct SuggestMoviesData
		if res, err := c.Suggest(context.Background(), in); err == nil {
			for {
				msg, err := res.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("suggest error: %v", err)
				}
				log.Printf("Suggest Response: %v", msg.GetTitle())
				currentMovie := SuggestMovie{msg.GetTitle()}
				mainStruct.Movies = append(mainStruct.Movies, currentMovie)
			}
			ctx.JSON(http.StatusOK, mainStruct)
		}
	})
	if err := g.Run(":8090"); err != nil {
		log.Fatalf("Failed to run client: %v", err)
	}
}
