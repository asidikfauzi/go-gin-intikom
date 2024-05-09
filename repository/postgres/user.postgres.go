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

func (u *User) GetAll(param model.ParamPaginate) (users []model.GetUser, count int64, err error) {

	query := u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("(name ILIKE ? OR email ILIKE ?)", "%"+param.Search+"%", "%"+param.Search+"%")

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	if err = query.Offset(param.Offset).
		Order(param.OrderBy + " " + param.Direction).
		Find(&users).Error; err != nil {
		return
	}

	if err = query.Count(&count).Error; err != nil {
		return
	}

	return
}

func (u *User) FindById(id int) (user model.Users, err error) {

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

func (u *User) IdExists(id int) bool {
	var user model.Users
	if err := u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&user).Error; err != nil {
		return false
	}

	return true
}

func (u *User) EmailExists(email string) bool {
	var user model.Users
	if err := u.DB.Model(&model.Users{}).
		Where("deleted_at IS NULL").
		Where("email = ?", email).
		First(&user).Error; err != nil {
		return false
	}

	return true
}

func (u *User) Create(user *model.Users) error {
	return u.DB.Create(user).Error
}

func (u *User) Update(user model.Users) error {
	query := u.DB.Model(&model.Users{}).
		Where("id = ?", user.ID)

	if user.Name != "" {
		query = query.UpdateColumn("name", user.Name)
	}

	if user.Email != "" {
		query = query.UpdateColumn("email", user.Email)
	}

	if user.Password != "" {
		query = query.UpdateColumn("password", user.Password)
	}

	query = query.UpdateColumn("updated_at", user.UpdatedAt)

	if err := query.Error; err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(user model.Users) error {
	if err := u.DB.Model(&model.Users{}).
		Where("id = ?", user.ID).
		UpdateColumn("deleted_at", user.DeletedAt).Error; err != nil {
		return err
	}

	return nil
}
