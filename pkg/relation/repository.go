package relation

type Repository interface{
	Close()
	CreatRelation(r *Relation) error
	GetProfileFollowing(f *[]int,profileId int) error
	GetProfileFollower(f *[]int,profileId int) error
}

func Close() {
	DB.Close()
}

func CreatRelation(r *Relation) error{
	err:=DB.Create(r).Error
	if err!=nil{
		return err
	}
	return nil
}

func GetProfileFollowing(f *[]int,profileId int) error{
	var r []Relation
	err:=DB.Where("profile_id = ?",profileId).Find(&r).Error
	if err!=nil{
		return err
	}
	for _,a := range r {
		*f=append(*f,a.FollowingId)
	}
	return nil
}

func GetProfileFollower(f *[]int,profileId int) error{
	var r []Relation
	err:=DB.Where("following_id = ?",profileId).Find(&r).Error
	if err!=nil{
		return err
	}
	for _,a := range r {
		*f=append(*f,a.ProfileId)
	}
	return nil
}