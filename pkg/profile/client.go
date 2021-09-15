package profile

import (
	"context"
	"gosocial/profile/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.ProfileServiceClient
}

func NewClient(url string)(*Client,error){
	cinn,err:=grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c:=pb.NewProfileServiceClient(cinn)
	return &Client{cinn,c},nil
}

func (c *Client) Close(){
	c.conn.Close()
}

func (c *Client) PostProfile(ctx context.Context,username string)(*Profile,error){
	p,err:=c.service.PostProfile(
		ctx,
		&pb.PostProfileRequest{UserName: username},
	)
	if err != nil {
		return nil, err
	}
	return &Profile{
		ID: int(p.Profile.Id),
		UserName: p.Profile.UserName,
	},nil
}

func (c *Client) GetProfile(ctx context.Context,id int)(*Profile,error){
	p,err:=c.service.GetProfile(ctx,&pb.GetProfileRequest{Id: uint64(id)})
	if err != nil {
		return nil, err
	}
	return &Profile{
		ID: int(p.Profile.Id),
		UserName: p.Profile.UserName,
	},nil
}

func (c *Client) GetAllProfiles(ctx context.Context)([]Profile,error){
	p,err:=c.service.GetAllProfiles(ctx,&pb.GetAllProfilesRequest{})
	if err != nil {
		return nil, err
	}
	profiles:=[]Profile{}
	for _,a := range p.Profile {
		profiles=append(profiles, Profile{
			ID: int(a.Id),
			UserName: a.UserName,
		})
	}
	return profiles,nil
}

