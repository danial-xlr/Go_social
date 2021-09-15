package feed

import (
	"context"
	"time"
)

type Service interface{
	CreatFeed(ctx context.Context,profileId int,posts []FeedPost)(*Feed,error)
	GetProfileFeed(ctx context.Context,profileId int)(*Feed,error)
}

type Feed struct {
	ID        	int `json:"id"`
	ProfileId 	int `json:"profile_id"`
	Posts 		[]FeedPost `json:"feed_post"`
}

type FeedPost struct {
	ID        int       `json:"id"`
	Body      string    `json:"body"`
	ProfileId int       `json:"profile_id"`
	CreatAt   time.Time `json:"creat_at"`
}

type feedService struct{
	repository Repository
}

func NewService(r Repository) Service{
	return &feedService{r}
}

func (t *feedService)CreatFeed(ctx context.Context,profileId int,posts []FeedPost)(*Feed,error){
	f:=&Feed{
		ProfileId: profileId,
		Posts: posts,
	}
	err:=t.repository.CreatFeed(f)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func (t *feedService)GetProfileFeed(ctx context.Context,profileId int)(*Feed,error){
	var f Feed
	err:=t.repository.GetProfileFeed(&f,profileId)
	if err != nil {
		return nil, err
	}
	return &f, nil
}

