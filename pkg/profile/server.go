//go:generate protoc .\profile.proto --go_out=plugins=grpc:.\pb
package profile

import (
	"net"
	"context"

	"gosocial/profile/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct{
	sevice Service
}

func ListenGRPC(s Service) error {
	lis,err:=net.Listen("tcp",":8000")
	if err != nil {
		return err
	}
	serv:=grpc.NewServer()	
	pb.RegisterProfileServiceServer(serv,&grpcServer{s})	
	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer) PostProfile(ctx context.Context,r *pb.PostProfileRequest ) (*pb.PostProfileResponse,error){
	p,err:=s.sevice.PostProfile(ctx,r.UserName)
	if err != nil {
		return nil, err
	}
	return &pb.PostProfileResponse{Profile: &pb.Profile{
		Id: uint64(p.ID),
		UserName: p.UserName,
	}},nil
}

func (s *grpcServer) GetProfile(ctx context.Context,r *pb.GetProfileRequest) (*pb.GetProfileResponse,error){
	p,err:=s.sevice.GetProfile(ctx,int(r.Id))
	if err != nil {
		return nil, err
	}
	return &pb.GetProfileResponse{Profile: &pb.Profile{
		Id: uint64(p.ID),
		UserName: p.UserName,
	}},nil
}

func (s *grpcServer) GetAllProfiles(ctx context.Context,r *pb.GetAllProfilesRequest) (*pb.GetAllProfilesResponse,error){
	p,err:=s.sevice.GetAllProfiles(ctx)
	if err != nil {
		return nil, err
	}
	profiles:=[]*pb.Profile{}
	for _,a:=range *p {
		profiles=append(profiles, &pb.Profile{
			Id: uint64(a.ID),
			UserName: a.UserName,
		},
	)
	}
	return &pb.GetAllProfilesResponse{Profile: profiles},nil
}
