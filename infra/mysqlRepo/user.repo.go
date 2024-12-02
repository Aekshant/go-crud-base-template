package infra

import (
	"goGinTemplate/domain"

	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAllUser() ([]domain.User, error)
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

func (u UserRepositoryImpl) FindAllUser() ([]domain.User, error) {
	var users []domain.User

	var err = u.db.Find(&users).Error
	if err != nil {
		log.Error("Got an error finding all users. Error: ", err)
		return nil, err
	}

	return users, nil
}
