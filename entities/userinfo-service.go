package entities

//UserInfoAtomicService .
type UserInfoAtomicService struct{}

//UserInfoService .
var UserInfoService = UserInfoAtomicService{}

// Save .
func (*UserInfoAtomicService) Save(u *UserInfo) error {
	// 当使用事务处理时，需要创建Session对象。在进行事物处理时，可以混用ORM方法和RAW方法
	session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	err := session.Begin()
	_, err = session.Insert(&u)
	if err != nil {
	  session.Rollback()
	  return err
	}
	// add Commit() after all actions
	err = session.Commit()
	if err != nil {
	  return err
	}
	return nil
}

// FindAll .
func (*UserInfoAtomicService) FindAll() []UserInfo {
	// Find方法的第一个参数为slice的指针或Map指针
	everyone := make([]UserInfo, 0)
	err := engine.Find(&everyone)
  checkErr(err)
	return everyone

	/*pEveryOne := make([]*Userinfo, 0)
	err := engine.Find(&pEveryOne)*/
}
// 通过调用engine.DBMetas()可以获取到数据库中所有的表，字段，索引的信息

// FindByID .
func (*UserInfoAtomicService) FindByID(id int) *UserInfo {
	// 传入一个主键字段的值，作为查询条件
	var user UserInfo
	engine.Id(id).Get(&user)
	// SELECT * FROM user Where id = id

	return &user


	/*// 根据Id来获得单条数据
	user := new(User)
	has, err := engine.Id(id).Get(user)*/


	/*// 直接执行一个SQL查询
	sql := "select * from user"
	results, err := engine.Query(sql)*/
}
