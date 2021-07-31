package main

import (
	"context"
	pb "github.com/xgSama/go_learning/grpc/src/route"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:5000")
	if err != nil {
		log.Fatalln(err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterRouteGuideServer(grpcServer, newServer())

	log.Fatalln(grpcServer.Serve(lis))

}

type routeGuideServer struct {
	pb.UnimplementedRouteGuideServer
}

func (s *routeGuideServer) GetFeature(ctx context.Context, point *pb.Point) (*pb.Feature, error) {
	return nil, nil
}

func (s *routeGuideServer) ListFeature(r *pb.Rectangle, d pb.RouteGuide_ListFeatureServer) error {
	return nil
}

func (s *routeGuideServer) RecordRoute(pb.RouteGuide_RecordRouteServer) error {
	return nil
}

func (s *routeGuideServer) Recommend(pb.RouteGuide_RecommendServer) error {
	return nil
}

func newServer() *routeGuideServer {
	return &routeGuideServer{}
}
