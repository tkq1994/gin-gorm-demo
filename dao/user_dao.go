package dao

var userDao *UserDao

type UserDao struct {
	BaseDao
}

func NewUserDao() *UserDao {
	if userDao == nil {
		return &UserDao{NewBaseDao()}
	}

	return userDao
}
