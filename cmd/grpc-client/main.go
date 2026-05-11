package main

import (
	"context"
	"log"
	"time"

	pb "mangahub/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient(
		"127.0.0.1:50051",
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
	)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	client := pb.NewMangaServiceClient(conn)

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	res, err := client.GetMangaList(
		ctx,
		&pb.Empty{},
	)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("=== Manga List ===")

	for _, manga := range res.Mangas {
		log.Println(
			manga.Id,
			manga.Title,
			manga.Author,
		)
	}
}