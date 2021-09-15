package relation

import "context"

type Service interface {
	CreatRelation(ctx context.Context,profileid int,followingid int)(*Relation,error)
	GetProfileFollowing(ctx context.Context,profileid int)(*[]int,error)
	GetProfileFollower(ctx context.Context,profileid int)(*[]int,error)
}

type Relation struct {
	ID          int `json:"id"`
	ProfileId   int `json:"profile_id"`
	FollowingId int `json:"following_id"`
}

type relationService struct {
	repository Repository
}

func NewService(r Repository) Service {
	return &relationService{r}
}

func (t *relationService) CreatRelation(ctx context.Context,profileid int,followingid int)(*Relation,error){
	r:=Relation{
		ProfileId: profileid,
		FollowingId: followingid,
	}
	err:=t.repository.CreatRelation(&r)
	if err!=nil{
		return nil,err
	}
	return &r,nil
}

func (t *relationService) GetProfileFollowing(ctx context.Context,profileid int)(*[]int,error){
	var f []int
	err:=t.repository.GetProfileFollowing(&f,profileid)
	if err!=nil{
		return nil,err
	}
	return &f,nil
}

func (t *relationService) GetProfileFollower(ctx context.Context,profileid int)(*[]int,error){
	var f []int
	err:=t.repository.GetProfileFollower(&f,profileid)
	if err!=nil{
		return nil,err
	}
	return &f,nil
}
