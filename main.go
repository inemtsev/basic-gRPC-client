package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	basketballPlayer "basic-gRPC-proto"
	"log"
)

func main() {
	router := gin.Default()

	rg := router.Group("api/v1/basketballPlayer")
	{
		rg.GET("/:id", fetchBasketballPlayer)
		// Add more routes here later
	}

	router.Run()
}

func fetchBasketballPlayer(c *gin.Context) {
	sAddress := "<yourServerIP>:50051"
	conn, e := grpc.Dial(sAddress, grpc.WithInsecure())
	if e != nil {
		log.Fatalf("Failed to connect to server %v", e)
	}
	defer conn.Close()

	client := basketballPlayer.NewPlayerServiceClient(conn)
	player, e := client.GetBasketballPlayer(context.Background(), &basketballPlayer.PlayerRequest{
		Id:                   c.Param("id"),
	})
	if e != nil {
		log.Fatalf("Failed to get player data: %v", e)
	}

	c.JSON(200, &player)
}