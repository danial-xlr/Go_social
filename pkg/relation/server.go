//go:generate protoc .\relation.proto --go_out=plugins=grpc:.\pb
package relation

import (
	"context"
	"errors"
	"fmt"
	"gosocial/profile"
	"gosocial/relation/pb"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	srvice        Service
	profileClient *profile.Client
}

func ListenGRPC(s Service,profileURL string,port int)error{
	profileClient,err:=profile.NewClient(profileURL)
	if err != nil {
		return err
	}
	lis,err:=net.Listen("tcp",fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		return err

	}
	serv:=grpc.NewServer()
	pb.RegisterRelationServiceServer(serv,&grpcServer{
		s,
		profileClient,
	})
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer)CreatRelation(ctx context.Context,r *pb.CreatRelationRequest)(*pb.CreatRelationResponse,error){
	_,err :=s.profileClient.GetProfile(ctx,int(r.ProfileId))
	if err != nil {
		log.Println("Error getting account: ", err)
		return nil, errors.New("account not found")
	}
	_,err = s.profileClient.GetProfile(ctx,int(r.FollowingId))
	if err != nil {
		log.Println("Error getting account: ", err)
		return nil, errors.New("account not found")
	}
	rel,err:=s.srvice.CreatRelation(ctx,int(r.ProfileId),int(r.FollowingId))
	if err != nil {
		return nil, err
	}
	return &pb.CreatRelationResponse{Relation: &pb.Relation{
		Id: uint64(rel.ID),
		ProfileId: uint64(rel.ProfileId),
		FollowingId: uint64(rel.FollowingId),
	}},nil
}

func (s *grpcServer)GetProfileFollower(ctx context.Context, r *pb.GetProfileFollowerRequest)(*pb.GetProfileFollowerResponse,error){
	rel,err:=s.srvice.GetProfileFollower(ctx,int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	f:= []uint64{}
	for _,a:=range *rel{
		f=append(f, uint64(a))
	}
	return &pb.GetProfileFollowerResponse{ProfileId: f},nil
}

func (s *grpcServer)GetProfileFollowing(ctx context.Context,r *pb.GetProfileFollowingRequest)(*pb.GetProfileFollowingResponse,error){
	rel,err:=s.srvice.GetProfileFollowing(ctx,int(r.ProfileId))
	if err != nil {
		return nil, err
	}
	f:= []uint64{}
	for _,a:=range *rel{
		f=append(f, uint64(a))
	}
	return &pb.GetProfileFollowingResponse{ProfileId: f},nil
}

