package mysqlrepo

import (
	"goGinTemplate/domain/dao"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]dao.User, error)
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

func UserRepositoryInit(db *gorm.DB) *UserRepositoryImpl {
	//db.AutoMigrate(&dao.User{})
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) FindAllUser() ([]dao.User, error) {
	var users []dao.User

	var err = u.db.Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all users. Error: ", err)
		return nil, err
	}

	return users, nil
}
