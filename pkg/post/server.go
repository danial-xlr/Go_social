//go:generate protoc .\post.proto --go_out=plugins=grpc:.\pb
package post

import (
	"context"
	"errors"
	"fmt"
	"gosocial/post/pb"
	"gosocial/profile"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service Service
	profileClient *profile.Client
}

func ListenGRPC(s Service, profileURL string, port int) error {
	profileClient, err := profile.NewClient(profileURL)
	if err != nil {
		return err
	}
	lis,err:=net.Listen("tcp",fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		return err

	}
	serv:=grpc.NewServer()
	pb.RegisterPostServiceServer(serv,&grpcServer{
		s,
		profileClient,
	})
	reflection.Register(serv)

	return serv.Serve(lis)
}

func (s *grpcServer)CreatPost(ctx context.Context,r *pb.CreatPostRequest)(*pb.CreatPostResponse,error){
	_,err:=s.profileClient.GetProfile(ctx,int(r.ProfileId))
	if err != nil {
		log.Println("Error getting account: ", err)
		return nil, errors.New("account not found")
	}
	p,err:=s.service.CreatPost(ctx,int(r.ProfileId),r.Body)
	if err != nil {
		return nil, err
	}
	return &pb.CreatPostResponse{Post: &pb.Post{
		Id: uint64(p.ID),
		ProfileId: uint64(p.ProfileId),
		Body: p.Body,
	}},nil

}

func (s *grpcServer)GetPostById(ctx context.Context,r *pb.GetPostByIdRequest)(*pb.GetPostByIdResponse,error){
	p,err:=s.service.GetPostById(ctx,int(r.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetPostByIdResponse{Post: &pb.Post{
		Id: uint64(p.ID),
		ProfileId: uint64(p.ProfileId),
		Body: p.Body,		
	}},nil
}

func (s *grpcServer)GetPostsByProfileId(ctx context.Context,r *pb.GetPostsByProfileIdRequest)(*pb.GetPostsByProfileIdResponse,error){
	p,err:=s.service.GetPostsByProfileId(ctx,int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	posts := []*pb.Post{}
	for _,a:=range *p {
		posts=append(posts, &pb.Post{
			Id: uint64(a.ID),
			ProfileId: uint64(a.ProfileId),
			Body: a.Body,
		})
	}
	return &pb.GetPostsByProfileIdResponse{Post: posts},nil
}