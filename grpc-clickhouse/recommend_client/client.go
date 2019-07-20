package main

import (
	"context"
	"fmt"
	"grpc-clickhouse/recommendpb"
	"io"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

type MoviesData struct {
	Movies []Movie `json:"Movie"`
}

type Movie struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
}

func main() {
	cc, err := grpc.Dial("localhost:50050", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer cc.Close()
	c := recommendpb.NewRecommendServiceClient(cc)
	clickStream(c)
}
func clickStream(c recommendpb.RecommendServiceClient) {
	fmt.Println("starting rpc stream")
	g := gin.Default()
	g.Use(cors.Default())

	g.GET("/recommend/:id", func(ctx *gin.Context) {
		id := ctx.Param("id")
		in := &recommendpb.RecommendRequest{
			Recommender: &recommendpb.Recommender{
				Movie: id,
			},
		}
		var mainStruct MoviesData
		if res, err := c.Recommend(context.Background(), in); err == nil {
			for {
				msg, err := res.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("recommend error: %v", err)
				}
				log.Printf("Response from GreetManyTimes: %v", msg.GetMovieTitle())
				currentMovie := Movie{msg.GetMovieId(), msg.GetMovieTitle()}
				mainStruct.Movies = append(mainStruct.Movies, currentMovie)
			}
			ctx.JSON(http.StatusOK, mainStruct)
		}
	})
	if err := g.Run(":8091"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
