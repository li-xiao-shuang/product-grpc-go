package main

import (
	"context"
	"fmt"
	"github.com/gofrs/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"net"
	pb "prodcuctinfo/proto"
)

// 用来实现product_info 的服务器
type server struct {
	productMap map[string]*pb.Product
}

// 添加一个产品
func (s *server) AddProduct(ctx context.Context, in *pb.Product) (*pb.ProductId, error) {
	fmt.Println("添加一个产品,产品信息：", in)
	out, err := uuid.NewV4()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error while generating Product ID", err)
	}
	in.Id = out.String()
	if s.productMap == nil {
		s.productMap = make(map[string]*pb.Product)
	}
	s.productMap[in.Id] = in
	return &pb.ProductId{Value: in.Id}, status.New(codes.OK, "").Err()
}

// 获取产品信息
func (s *server) GetProduct(ctx context.Context, id *pb.ProductId) (*pb.Product, error) {
	fmt.Println("获取产品信息，产品id：", id)
	values, error := s.productMap[id.Value]
	if error {
		return values, status.New(codes.OK, "").Err()
	}
	return nil, status.Errorf(codes.NotFound, "Product does not exist", id.Value)
}

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listend:%v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProductInfoServer(s, &server{})
	log.Printf("starting grpc listener on port " + port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve :%v", err)
	}
}
