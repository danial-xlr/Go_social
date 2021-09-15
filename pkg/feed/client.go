package feed

import (
	"context"
	"gosocial/feed/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	Service pb.FeedServiceClient
}

func NewClient(url string)(*Client,error){
	cinn,err:=grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c:=pb.NewFeedServiceClient(cinn)
	return &Client{cinn,c},nil
}

func (c *Client) Close(){
	c.conn.Close()
}

func (c *Client)CreatFeed(ctx context.Context,profile_id int)(*Feed,error){
	feed,err:=c.Service.CreatFeed(ctx,&pb.CreatFeedRequest{
		ProfileId: uint64(profile_id),
	})
	if err != nil {
		return nil, err
	}
	protoPost:= []FeedPost{}
	for _,a:=range feed.Feed.FeedPost{
		protoPost=append(protoPost, FeedPost{
			ID: int(a.Id),
			ProfileId: int(a.ProfileId),
			Body: a.Body,
			//CreatAt: a.CreatAt,
		})
	}
	return &Feed{
		ID: int(feed.Feed.Id),
		ProfileId: int(feed.Feed.ProfileId),
		Posts: protoPost,
	},nil
}

func (c *Client)GetProfileFeed(ctx context.Context,profile_id int)(*Feed,error){
	feed,err:=c.Service.GetProfileFeed(ctx,&pb.GetProfileFeedRequest{
		ProfileId: uint64(profile_id),
	})
	if err != nil {
		return nil, err
	}
	protoPost:= []FeedPost{}
	for _,a:=range feed.Feed.FeedPost{
		protoPost=append(protoPost, FeedPost{
			ID: int(a.Id),
			ProfileId: int(a.ProfileId),
			Body: a.Body,
			//CreatAt: a.CreatAt,
		})
	}
	return &Feed{
		ID: int(feed.Feed.Id),
		ProfileId: int(feed.Feed.ProfileId),
		Posts: protoPost,
	},nil
}

