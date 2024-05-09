package postgres

import (
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"gorm.io/gorm"
)

type User struct {
	DB *gorm.DB
}

func NewUserPostgres(conn *gorm.DB) domain.UserPostgres {
	return &User{
		DB: conn,
	}
}

func (u *User) GetAll(limit, offset int, orderBy, search string) (users []model.GetUser, count int64, err error) {

	if err = u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("(id ILIKE ? OR name ILIKE ? OR email ILIKE ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Limit(limit).
		Offset(offset).
		Order(orderBy).
		Find(&users).Error; err != nil {
		return
	}

	if err = u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("(id ILIKE ? OR name ILIKE ? OR email ILIKE ?)", "%"+search+"%", "%"+search+"%", "%"+search+"%").
		Count(&count).Error; err != nil {
		return
	}

	return
}

func (u *User) FindById(id int) (user model.GetUser, err error) {

	if err = u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return
	}

	return
}

func (u *User) FindByEmail(email string) (user model.Users, err error) {
	if err = u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return
	}

	return
}
