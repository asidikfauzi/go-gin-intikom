package postgres

import (
	"asidikfauzi/go-gin-intikom/domain"
	"asidikfauzi/go-gin-intikom/model"
	"gorm.io/gorm"
)

type Task struct {
	DB     *gorm.DB
	userPg domain.UserPostgres
}

func NewTaskPostgres(conn *gorm.DB, up domain.UserPostgres) domain.TaskPostgres {
	return &Task{
		DB:     conn,
		userPg: up,
	}
}

func (u *Task) GetAll(param model.ParamPaginate) (tasks []model.GetTask, count int64, err error) {

	query := u.DB.Model(&model.Tasks{}).
		Where("deleted_at IS NULL").
		Where("(title ILIKE ? OR description ILIKE ? OR status ILIKE ?)", "%"+param.Search+"%", "%"+param.Search+"%", "%"+param.Search+"%")

	if err = query.Count(&count).Error; err != nil {
		return
	}

	if param.Limit > 0 {
		query = query.Limit(param.Limit)
	}

	if err = query.Offset(param.Offset).
		Order(param.OrderBy + " " + param.Direction).
		Find(&tasks).Error; err != nil {
		return
	}

	for i := range tasks {
		user, errUser := u.userPg.FindById(int(tasks[i].UserID))
		if errUser != nil {
			return
		}

		tasks[i].User = model.GetUser{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		}
	}

	return
}

func (u *Task) FindById(id int) (task model.Tasks, err error) {

	if err = u.DB.Model(&model.Tasks{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&task).Error; err != nil {
		return
	}

	return
}

func (u *Task) IdExists(id int) bool {
	var task model.Tasks
	if err := u.DB.Model(&model.Tasks{}).
		Where("deleted_at IS NULL").
		Where("id = ?", id).
		First(&task).Error; err != nil {
		return false
	}

	return true
}

func (u *Task) TitleExists(title string) bool {
	var task model.Tasks
	if err := u.DB.Model(&model.Tasks{}).
		Where("deleted_at IS NULL").
		Where("title = ?", title).
		First(&task).Error; err != nil {
		return false
	}

	return true
}

func (u *Task) Create(task *model.Tasks) error {
	return u.DB.Create(task).Error
}

func (u *Task) Update(task model.Tasks) error {
	query := u.DB.Model(&model.Tasks{}).
		Where("id = ?", task.ID)

	if task.UserID != 0 {
		query = query.UpdateColumn("user_id", task.UserID)
	}

	if task.Title != "" {
		query = query.UpdateColumn("title", task.Title)
	}

	if task.Description != "" {
		query = query.UpdateColumn("description", task.Description)
	}

	if task.Status != "" {
		query = query.UpdateColumn("status", task.Status)
	}

	query = query.UpdateColumn("updated_at", task.UpdatedAt)

	if err := query.Error; err != nil {
		return err
	}

	return nil
}

func (u *Task) Delete(task model.Tasks) error {
	if err := u.DB.Model(&model.Tasks{}).
		Where("id = ?", task.ID).
		UpdateColumn("deleted_at", task.DeletedAt).Error; err != nil {
		return err
	}

	return nil
}
