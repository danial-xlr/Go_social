package feed

type Repository interface{
	Close()
	CreatFeed(f *Feed) error
	GetProfileFeed(f *Feed,profileId int) error
}

func Close() {
	DB.Close()
}

func CreatFeed(f *Feed) error{
	err:=DB.Create(f).Error
	if err!=nil{
		return err
	}
	return nil
}

func GetProfileFeed(f *Feed,profileId int) error{
	err:=DB.Where("profile_id = ?",profileId).Find(&f).Error
	if err!=nil{
		return err
	}
	return nil
}