package relation

import (
	"context"
	"gosocial/relation/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	Service pb.RelationServiceClient
}

func NewClient(url string)(*Client,error){
	cinn,err:=grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c:=pb.NewRelationServiceClient(cinn)
	return &Client{cinn,c},nil
}

func (c *Client) Close(){
	c.conn.Close()
}

func (c *Client) CreatRelation(ctx context.Context,profile_id int ,following_id int)(*Relation,error){
	rel,err:=c.Service.CreatRelation(ctx,&pb.CreatRelationRequest{
		ProfileId: uint64(profile_id),
		FollowingId: uint64(following_id),
	})
	if err != nil {
		return nil, err
	}
	return &Relation{
		ID: int(rel.Relation.Id),
		ProfileId: int(rel.Relation.ProfileId),
		FollowingId: int(rel.Relation.FollowingId),
	},nil
}

func (c *Client) GetProfileFollower(ctx context.Context,profile_id int)([]uint64,error){
	rel,err:=c.Service.GetProfileFollower(ctx,&pb.GetProfileFollowerRequest{
		ProfileId: uint64(profile_id),
	})
	if err != nil {
		return nil, err
	}
	 f:= []uint64{}
	for _,a:=range rel.ProfileId {
		f=append(f, uint64(a))
	} 
	return f,nil
}

func (c *Client) GetProfileFollowing(ctx context.Context,profile_id int)([]uint64,error){
	rel,err:=c.Service.GetProfileFollowing(ctx,&pb.GetProfileFollowingRequest{
		ProfileId: uint64(profile_id),
	})
	if err != nil {
		return nil, err
	}
	 f:= []uint64{}
	for _,a:=range rel.ProfileId {
		f=append(f, uint64(a))
	} 
	return f,nil
}

