package profile

import "context"

type Service interface {
	PostProfile(c context.Context,Uname string)(*Profile,error)
	GetProfile(c context.Context, id int)(*Profile,error)
	GetAllProfiles(c context.Context)(*[]Profile,error)
}

type Profile struct {
	ID       int    `json:"id"`
	UserName string `json:"user_name"`
}

type profileService struct{
	repo Repository
}

func NewService(r Repository) Service {
	return &profileService{r}
}

func (r *profileService) PostProfile(c context.Context,Uname string)(*Profile,error){
	p :=&Profile{
		UserName: Uname,		
	}
	e:= r.repo.CreatProfile(p)
	return p,e
}

func (r *profileService)GetProfile(c context.Context, id int)(*Profile,error){
	var p Profile
	e:=r.repo.GetProfileByid(&p,id)
	return &p,e
}

func (r *profileService)GetAllProfiles(c context.Context)(*[]Profile,error){
	var p []Profile
	e:=r.repo.ReadAllProfiles(&p)
	return &p,e
}