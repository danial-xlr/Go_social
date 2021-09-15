package post

import (
	"context"
	"time"
)

type Service interface{
	CreatPost(ctx context.Context,profileid int,body string)(*Post,error)
	GetPostById(ctx context.Context,postId int)(*Post,error)
	GetPostsByProfileId(ctx context.Context,profileId int)(*[]Post,error)
}

type Post struct {
	ID        int    `json:"id"`
	Body      string `json:"body"`
	ProfileId int    `json:"profile_id"`
	CreatAt   time.Time `json:"creat_at"`
}

type postService struct{
	repository Repository
}

func NewService(r Repository) Service {
	return &postService{r}
}

func (t *postService)CreatPost(ctx context.Context,profileid int,body string)(*Post,error){
	p:=Post{
		Body: body,
		ProfileId: profileid,
		CreatAt: time.Now(),
	}
	err:=t.repository.CreatPost(&p)
	if err!=nil{
		return nil,err
	}
	return &p,nil
}

func (t *postService)GetPostById(ctx context.Context,id int)(*Post,error){
	var p Post
	err:=t.repository.GetPostById(&p,id)
	if err!=nil{
		return nil,err
	}
	return &p,nil

}

func (t *postService) GetPostsByProfileId(ctx context.Context,profileId int)(*[]Post,error){
	var p []Post
	err:=t.repository.GetPostsByProfileId(&p,profileId)
	if err!=nil{
		return nil,err
	}
	return &p,nil
}