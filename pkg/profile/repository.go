package profile

type Repository interface{
	Close()
	CreatProfile(p *Profile)(err error)
	GetProfileByid(p *Profile ,id int)(err error)
	ReadAllProfiles(p *[]Profile)(err error)
}

func Close(){
	DB.Close()
}

func CreatProfile(p *Profile)(err error){
	err=DB.Create(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func GetProfileByid(p *Profile ,id int)(err error){	
	err=DB.Where("id = ?",id).First(p).Error
	if err!=nil{
		return err
	}
	return nil
}

func ReadAllProfiles(p *[]Profile)(err error){
	err=DB.Find(p).Error
	if err!=nil{
		return err
	}
	return nil
}