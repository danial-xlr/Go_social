package post


type Repository interface {
	Close()
	CreatPost(p *Post)error
	GetPostById(p *Post,id int)error
	GetPostsByProfileId(p *[]Post,profileId int)error
}

func Close() {
	DB.Close()
}

func CreatPost(p *Post) error {
	err:=DB.Create(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func GetPostById(p *Post,id int)error{
	err:=DB.Where("id = ?",id).First(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func GetPostsByProfileId(p *[]Post,profileId int)error{
	err:=DB.Where("profile_id = ?",profileId).Find(p).Error
	if err!=nil{
		return err
	}
	return nil
}