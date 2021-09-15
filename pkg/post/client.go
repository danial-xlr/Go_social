package post

import (
	"context"
	"gosocial/post/pb"

	"google.golang.org/grpc"
)

type Client struct {
	conn *grpc.ClientConn
	service pb.PostServiceClient
}

func NewClient(url string)(*Client,error){
	cinn,err:=grpc.Dial(url,grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c:=pb.NewPostServiceClient(cinn)
	return &Client{cinn,c},nil
}

func (c *Client) Close(){
	c.conn.Close()
}

func (c *Client) CreatPost(ctx context.Context,body string,profile_id int)(*Post,error){
	p,err:=c.service.CreatPost(ctx,&pb.CreatPostRequest{
		Body: body,
		ProfileId: uint64(profile_id),
	})
	if err != nil {
		return nil, err
	}
	return &Post{
		ID: int(p.Post.Id),
		Body: p.Post.Body,
		ProfileId: int(p.Post.Id),
	},nil
}

func (c *Client) GetPostById(ctx context.Context,id int)(*Post,error){
	p,err:=c.service.GetPostById(ctx,&pb.GetPostByIdRequest{Id: uint64(id)})
	if err != nil {
		return nil, err
	}
	return &Post{
		ID: int(p.Post.Id),
		Body: p.Post.Body,
		ProfileId: int(p.Post.Id),
	},nil
}

func (c *Client) GetPostsByProfileId(ctx context.Context,profileId int)([]Post,error){
	p,err:=c.service.GetPostsByProfileId(ctx,&pb.GetPostsByProfileIdRequest{
		ProfileId: uint64(profileId),
	})
	if err != nil {
		return nil, err
	}
	posts := []Post{}
	for _,a:=range p.Post {
		posts=append(posts, Post{
			ID: int(a.Id),
			ProfileId: int(a.ProfileId),
			Body: a.Body,
		})
	}
	return posts,nil
}

