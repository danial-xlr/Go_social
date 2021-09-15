//go:generate protoc .\feed.proto --go_out=plugins=grpc:.\pb
package feed

import (
	"context"
	"errors"
	"fmt"
	"gosocial/feed/pb"
	"gosocial/post"
	"gosocial/profile"
	"gosocial/relation"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type grpcServer struct {
	service        Service
	profileClient *profile.Client
	relationClient *relation.Client
	postClient *post.Client
}

func ListenGRPC(s Service,profileURL string,relationURL string,postURL string,port int)error{
	profileClient,err:=profile.NewClient(profileURL)
	if err != nil {
		return err
	}
	relationClient,err:=relation.NewClient(relationURL)
	if err != nil {
		profileClient.Close()
		return err
	}
	postClient,err:=post.NewClient(postURL)
	if err != nil {
		profileClient.Close()
		relationClient.Close()
		return err
	}
	lis,err:=net.Listen("tcp",fmt.Sprintf(":%d", port))
	if err != nil {
		profileClient.Close()
		relationClient.Close()
		postClient.Close()
		return err
	}
	serv:=grpc.NewServer()
	pb.RegisterFeedServiceServer(serv,&grpcServer{
		s,
		profileClient,
		relationClient,
		postClient,
	})

	reflection.Register(serv)
	return serv.Serve(lis)
}

func (s *grpcServer)CreatFeed(ctx context.Context,f *pb.CreatFeedRequest)(*pb.CreatFeedResponse,error){
	_,err :=s.profileClient.GetProfile(ctx,int(f.ProfileId))
	if err != nil {
		log.Println("Error getting account: ", err)
		return nil, errors.New("account not found")
	}
	following,err:=s.relationClient.GetProfileFollowing(ctx,int(f.ProfileId))
	if err != nil {
		return nil, err
	}
	feedposts:= []FeedPost{}
	for _,a:=range following{
		posts,_:=s.postClient.GetPostsByProfileId(ctx,int(a))
		for _,p := range posts{
			feedposts = append(feedposts, FeedPost{
				ID: p.ID,
				ProfileId: p.ProfileId,
				Body: p.Body,
				CreatAt: p.CreatAt,
			})
		}
	}
	feed,err:=s.service.CreatFeed(ctx,int(f.ProfileId),feedposts)
	if err != nil {
		return nil, err
	}
	protoPosts:= []*pb.Feed_FeedPost{}
	for _,po:=range feed.Posts{
		protoPosts=append(protoPosts, &pb.Feed_FeedPost{
			Id: uint64(po.ID),
			ProfileId: uint64(po.ProfileId),
			Body: po.Body,
			//CreatAt: po.CreatAt,
		})
	}
	return &pb.CreatFeedResponse{Feed: &pb.Feed{
		Id: uint64(feed.ID),
		ProfileId: uint64(feed.ProfileId),
		FeedPost: protoPosts,
	}},nil
}

func (s *grpcServer)GetProfileFeed(ctx context.Context,f *pb.GetProfileFeedRequest)(*pb.GetProfileFeedResponse,error){
	feed,err:=s.service.GetProfileFeed(ctx,int(f.ProfileId))
	if err != nil {
		return nil, err
	}
	protoPosts:= []*pb.Feed_FeedPost{}
	for _,po:=range feed.Posts{
		protoPosts=append(protoPosts, &pb.Feed_FeedPost{
			Id: uint64(po.ID),
			ProfileId: uint64(po.ProfileId),
			Body: po.Body,
			//CreatAt: po.CreatAt,
		})
	}
	return &pb.GetProfileFeedResponse{Feed: &pb.Feed{
		Id: uint64(feed.ID),
		ProfileId: uint64(feed.ProfileId),
		FeedPost: protoPosts,
	}},nil
}