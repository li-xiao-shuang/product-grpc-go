package main

import (
	"context"
	"google.golang.org/grpc"
	"log"
	pb "prodcuctinfo/proto"
	"time"
)

const address = "localhost:50051"

func main() {
	// 添加产品
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect:%v", err)
	}
	defer conn.Close()
	client := pb.NewProductInfoClient(conn)

	name := "苹果12"
	description := "非常好用"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.AddProduct(ctx, &pb.Product{Name: name, Description: description})
	if err != nil {
		log.Fatalf("could not add product: %v", err)
	}
	log.Printf("产品id : %s 添加成功", r.Value)

	// 获取产品
	resp,err := client.GetProduct(ctx,r)
	if err != nil {
		log.Fatalf("could not get product: %v", err)
	}
	log.Printf("获取产品信息 : %s ", resp.String())
}
