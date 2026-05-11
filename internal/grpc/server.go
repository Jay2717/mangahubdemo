package grpc

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "mangahub/proto"
)

type MangaServer struct {
	pb.UnimplementedMangaServiceServer
}

func (s *MangaServer) GetMangaList(
	ctx context.Context,
	req *pb.Empty,
) (*pb.MangaList, error) {

	mangas := []*pb.Manga{
		{
			Id:     "blue-box",
			Title:  "Blue Box",
			Author: "Kouji Miura",
		},
		{
			Id:     "oshi-no-koi",
			Title:  "Oshi no Ko",
			Author: "Aka Akasaka",
		},
	}

	return &pb.MangaList{
		Mangas: mangas,
	}, nil
}

func StartGRPCServer() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterMangaServiceServer(
		grpcServer,
		&MangaServer{},
	)

	log.Println("gRPC Server running on :50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
}